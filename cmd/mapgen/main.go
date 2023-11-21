package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/aquilax/go-perlin"
	"github.com/nfnt/resize"
)

const (
	// Define the image dimensions
	imageWidth  = 1920
	imageHeight = 1080

	quantizeColors  = 64
	oceanRangeStart = 0
	oceanRangeEnd   = 130
)

func main() {
	// Create a Perlin noise generator
	p := perlin.NewPerlin(6.312, 5.887, 8, time.Now().UnixNano())

	// Increase the dimensions for generating noise beyond the image borders
	noiseWidth := imageWidth
	noiseHeight := imageHeight

	// Create a new image
	img := image.NewPaletted(image.Rect(0, 0, imageWidth, imageHeight), color.Palette{})

	for i := 0; i < quantizeColors; i++ {
		val := 255 - uint8((float64(i)/quantizeColors-1)*255)

		r := val
		g := r
		b := r

		if val > oceanRangeStart && val < oceanRangeEnd {
			r = 0
			g = 0
			b = uint8(0)
		}

		img.Palette = append(img.Palette, color.RGBA{r, g, b, 255})
	}

	// Create a new Perlin noise image
	noiseImage := image.NewRGBA(image.Rect(0, 0, noiseWidth, noiseHeight))

	// Generate Perlin noise values for the extended dimensions

	for y := 0; y < noiseHeight; y++ {
		for x := 0; x < noiseWidth; x++ {
			n := p.Noise2D(float64(x)/float64(noiseWidth), float64(y)/float64(noiseHeight))

			index := int((n + 1.0) * ((quantizeColors - 1) / 2))
			noiseImage.Set(x, y, img.Palette[index])
		}
	}

	//rl.SetTargetFPS(30)
	//rl.InitWindow(imageWidth, imageHeight, "mapgen")
	//
	//raylibImage := rl.NewImageFromImage(noiseImage)
	//texture := rl.LoadTextureFromImage(raylibImage)
	//
	//for !rl.WindowShouldClose() {
	//	rl.BeginDrawing()
	//	rl.DrawTexture(texture, 0, 0, color.RGBA{255, 255, 255, 255})
	//	rl.EndDrawing()
	//}

	if err := createPNG("quantized_perlin_noise.png", noiseImage); err != nil {
		panic(err)
	}

	fmt.Println(ImageToASCII(noiseImage, 90))
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

// Convert an image.Image to ASCII art with the specified width
func ImageToASCII(img image.Image, width int) string {
	img = resizeImage(img, width)

	ascii := ""
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			brightness := float64(r+g+b) / (3 * 0xffff) // Normalize to [0, 1]

			// Map brightness to ASCII characters
			charIndex := int(brightness * float64(len(asciiGradient)-1))
			ascii += string(asciiGradient[charIndex])
		}
		if y != img.Bounds().Dy()-1 {
			ascii += "\n" // Add newline at the end of each row
		}
	}

	locations := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N'}
	for i := 0; i < len(locations); i++ {
		ascii = InsertCharacterInASCII(ascii, locations[i], i == 10)
	}

	ascii = ColorizeASCII(ascii)

	return ascii
}

// Resize an image to the specified width while maintaining aspect ratio
func resizeImage(img image.Image, width int) image.Image {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	aspectRatio := float64(dx) / float64(dy)
	newHeight := int(float64(width)/aspectRatio) / 2

	resizedImg := resize.Resize(uint(width), uint(newHeight), img, resize.Lanczos3)

	return resizedImg
}

var asciiGradient = []byte(" .,_-:;=+*#%@")

// InsertCharacterInASCII randomly inserts a character in the ASCII image
func InsertCharacterInASCII(asciiImage string, character rune, adjacentToEmpty bool) string {
	lines := strings.Split(asciiImage, "\n")
	numRows := len(lines)
	numCols := len(lines[0])

	if numRows == 0 || numCols == 0 {
		return asciiImage
	}

	// Create a copy of the ASCII image
	modifiedImage := make([]string, numRows)
	copy(modifiedImage, lines)

	// Find available positions (empty tiles)
	availablePositions := []struct{ row, col int }{}
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if string(modifiedImage[row][col]) != " " {
				if adjacentToEmpty {
					// Check if there"s an adjacent empty tile
					if isAdjacentEmpty(modifiedImage, row, col) {
						availablePositions = append(availablePositions, struct{ row, col int }{row, col})
					}
				} else {
					availablePositions = append(availablePositions, struct{ row, col int }{row, col})
				}
			}
		}
	}

	if len(availablePositions) == 0 {
		// No available positions to insert the character
		return asciiImage
	}

	// Randomly choose a position
	randomIndex := rand.Intn(len(availablePositions))
	position := availablePositions[randomIndex]

	// Insert the character at the chosen position
	modifiedImage[position.row] = modifiedImage[position.row][:position.col] +
		string(character) +
		modifiedImage[position.row][position.col+1:]

	// Reconstruct the modified ASCII image
	modifiedASCII := strings.Join(modifiedImage, "\n")

	return modifiedASCII
}

// Check if an adjacent tile is empty
func isAdjacentEmpty(asciiImage []string, row, col int) bool {
	numRows := len(asciiImage)
	numCols := len(asciiImage[0])

	// Check all adjacent positions
	positions := [][2]int{
		{row - 1, col}, // Above
		{row + 1, col}, // Below
		{row, col - 1}, // Left
		{row, col + 1}, // Right
	}

	for _, pos := range positions {
		r, c := pos[0], pos[1]
		if r >= 0 && r < numRows && c >= 0 && c < numCols && asciiImage[r][c] == ' ' {
			return true
		}
	}

	return false
}

// ColorizeASCII applies colors to the ASCII image with RGB foreground and background colors
func ColorizeASCII(asciiImage string) string {
	lines := strings.Split(asciiImage, "\n")
	numRows := len(lines)
	numCols := len(lines[0])

	if numRows == 0 || numCols == 0 {
		return asciiImage
	}

	// Create a copy of the ASCII image
	coloredImage := make([]string, numRows)
	copy(coloredImage, lines)

	for row := 0; row < numRows; row++ {
		newLine := ""
		for col := 0; col < numCols; col++ {
			char := string(coloredImage[row][col])

			switch char {
			case " ":
				// Dark blue background
				newLine += fmt.Sprintf("\x1b[48;2;32;32;64m%s\x1b[0m", char)
			case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N":
				// Red foreground with yellow background
				newLine += fmt.Sprintf("\x1b[38;2;255;0;0m\x1b[48;2;255;255;0m%s\x1b[0m", char)
			default:
				// Dark green background with light green foreground
				newLine += fmt.Sprintf("\x1b[38;2;173;255;47m\x1b[48;2;0;64;0m%s\x1b[0m", char)
			}
		}
		coloredImage[row] = newLine
	}

	// Reconstruct the colored ASCII image
	coloredASCII := strings.Join(coloredImage, "\n")

	return coloredASCII
}
