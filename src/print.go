package asciify

import (
	"fmt"
	"image"
	"image/color"
	"sort"
)

type tile struct {
	x      int
	y      int
	width  int
	height int
}

func scaleToSteps(lineHeight int, imageHeight int) int {
	return imageHeight / lineHeight
}

func sampleMean(img image.Image, tile tile) color.Gray {
	total := 0
	for y := tile.y; y < tile.y+tile.height; y++ {
		for x := tile.x; x < tile.x+tile.width; x++ {
			total += int(colorToGray(img.At(x, y)).Y)
		}
	}
	return color.Gray{uint8(total / (tile.width * tile.height))}
}

func sampleMedian(img image.Image, tile tile) color.Gray {
	values := []uint8{}
	for y := tile.y; y < tile.y+tile.height; y++ {
		for x := tile.x; x < tile.x+tile.width; x++ {
			values = append(values, uint8(colorToGray(img.At(x, y)).Y))
		}
	}
	less := func(i int, j int) bool {
		if i < j {
			return true
		}
		return false
	}
	sort.Slice(values, less)
	return color.Gray{values[len(values)/2]}
}

func sampleMid(img image.Image, tile tile) color.Gray {
	return colorToGray(img.At(tile.x+tile.width/2, tile.y+tile.height/2))
}

func sampleTopLeft(img image.Image, tile tile) color.Gray {
	return colorToGray(img.At(tile.x, tile.y))
}

func printImage(img image.Image, characterSet []rune, scaleSteps int) {
	stringImage := ""

	tileHeight := scaleSteps
	tileWidth := scaleSteps / 2

	for y := img.Bounds().Min.Y; y+tileHeight < img.Bounds().Max.Y; y += tileHeight {
		for x := img.Bounds().Min.X; x+tileWidth < img.Bounds().Max.X; x += tileWidth {
			// c := colorToGray(img.At(x, y))
			c := sampleMid(img, tile{x, y, tileWidth, tileHeight})
			stringImage += grayToChar(c, characterSet)
		}
		stringImage += "\n"
	}

	fmt.Print(stringImage)
}
