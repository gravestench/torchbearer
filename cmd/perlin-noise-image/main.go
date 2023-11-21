package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/aquilax/go-perlin"
)

func main() {
	// Create a Perlin noise generator
	p := perlin.NewPerlin(1, 1, 2, 0)

	// Define the image dimensions
	width := 512
	height := 512

	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Iterate over each pixel and set its color based on Perlin noise
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Generate Perlin noise value at this point
			n := p.Noise2D(float64(x)/float64(width), float64(y)/float64(height))

			// Map the Perlin noise value to an RGB color
			// You can adjust the mapping to get different effects
			r := uint8((n + 1.0) * 127.5)
			g := uint8((n + 1.0) * 127.5)
			b := uint8((n + 1.0) * 127.5)

			// Set the pixel color
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Create a new PNG file to save the image
	outputFile, err := os.Create("perlin_noise.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Encode and save the image as a PNG
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}
}
