package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	"github.com/aquilax/go-perlin"
)

const (
	// Define the image dimensions
	imageWidth  = 1920 >> 1
	imageHeight = 1080 >> 1

	quantizeColors  = 128
	oceanRangeStart = 0.0
	oceanRangeEnd   = 0.45
)

func main() {
	seed := time.Now().UnixNano()
	fmt.Printf("Seed: %v\r\n", seed)

	// Create a Perlin noise generator
	p := perlin.NewPerlin(6.312, 6.887, 12, seed)

	// Increase the dimensions for generating noise beyond the image borders
	noiseWidth := imageWidth * 2
	noiseHeight := imageHeight * 2

	const (
		shmooRange = 0.35
		shmooStep  = 0.0125
	)

	for endOcean := oceanRangeEnd - shmooRange; endOcean < oceanRangeEnd+shmooRange; endOcean += shmooStep {
		// Create a new image
		img := image.NewPaletted(image.Rect(0, 0, imageWidth, imageHeight), color.Palette{})

		for i := 0; i < quantizeColors; i++ {
			val := uint8((float64(i)/quantizeColors - 1) * 255)

			r := uint8(0)
			g := val - 32
			b := uint8(24) + (val >> 2)

			img.Palette = append(img.Palette, color.RGBA{r, g, b, 255})
		}

		// Create a new Perlin noise image
		noiseImage := image.NewRGBA(image.Rect(0, 0, noiseWidth, noiseHeight))

		// Generate Perlin noise values for the extended dimensions
		for y := 0; y < noiseHeight; y++ {
			for x := 0; x < noiseWidth; x++ {
				n := p.Noise2D(float64(x)/float64(noiseWidth), float64(y)/float64(noiseHeight))
				normalized := (n + 1.0) / 2
				index := int(normalized * (quantizeColors - 1))

				pixel := img.Palette[index]

				if normalized >= oceanRangeStart && normalized <= endOcean {
					pixel = color.RGBA{32, 32, 64, 255}
				}

				noiseImage.Set(x, y, pixel)
			}
		}

		// Find the largest contiguous section of non-ocean pixels
		largestNonOceanSection := findLargestNonOceanSection(noiseImage)

		if err := createPNG(fmt.Sprintf("largest_non_ocean_section_%f.png", endOcean), largestNonOceanSection); err != nil {
			panic(err)
		}
	}
}

func createPNG(name string, source *image.RGBA) error {
	// Create a new PNG file to save the image
	outputFile, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("creating image file: %v", err)
	}
	defer outputFile.Close()

	// Encode and save the image as a PNG
	err = png.Encode(outputFile, source)
	if err != nil {
		return fmt.Errorf("png encoding image: %v", err)
	}

	return nil
}

func findLargestNonOceanSection(img *image.RGBA) *image.RGBA {
	visited := make([][]bool, img.Bounds().Dy())
	for i := range visited {
		visited[i] = make([]bool, img.Bounds().Dx())
	}

	var largestSectionSize int
	var largestSection *image.RGBA

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			if !visited[y][x] {
				size, section := floodFill(img, x, y, visited)
				if size > largestSectionSize {
					largestSectionSize = size
					largestSection = section
				}
			}
		}
	}

	return largestSection
}

func floodFill(img *image.RGBA, x, y int, visited [][]bool) (int, *image.RGBA) {
	queue := []image.Point{{x, y}}
	visited[y][x] = true
	bounds := img.Bounds()
	section := image.NewRGBA(bounds)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		section.Set(p.X, p.Y, img.At(p.X, p.Y))

		for _, neighbor := range []image.Point{
			{p.X - 1, p.Y},
			{p.X + 1, p.Y},
			{p.X, p.Y - 1},
			{p.X, p.Y + 1},
		} {
			if neighbor.In(bounds) && !visited[neighbor.Y][neighbor.X] {
				nR, nG, nB, nA := img.At(neighbor.X, neighbor.Y).RGBA()
				if nR != 32 || nG != 32 || nB != 64 || nA != 255 {
					queue = append(queue, neighbor)
					visited[neighbor.Y][neighbor.X] = true
				}
			}
		}
	}

	return bounds.Dx() * bounds.Dy(), section
}
