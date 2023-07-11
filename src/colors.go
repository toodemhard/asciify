package asciify

import (
	"image/color"
	"math"
)

func squash(value float64) float64 {
	return math.Pow(value, 3) / 255 / 255
}

func colorToGray(pixel color.Color) color.Gray {
	// https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
	// r, g, b, _ := pixel.RGBA()
	// lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)

	return color.GrayModel.Convert(pixel).(color.Gray)
	// return color.Gray{uint8(lum / 256)}
}

func grayToChar(c color.Gray, characterSet []rune) string {
	levels := float64(len(characterSet))
	level := float64(c.Y) / 255.0 * levels
	if level == levels {
		level--
	}
	return string(characterSet[int(level)])
}
