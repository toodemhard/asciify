package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

// func image

//

func main() {
	// fmt.Println(os.Args[1])
	fileName := os.Args[1]
	reader, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	// _, imageType, err := image.Decode(reader)

	// if err != nil {
	// 	fmt.Println(fileName)
	// 	log.Fatal(err)
	// }

	img, err := png.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
	// fmt.Println(imageType)
	// image.Decode()
	// const width, height = 256, 256
	// img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		img.Set(x, y, color.NRGBA{
	// 			R: uint8((x + y) & 255),
	// 			G: uint8((x + y) << 1 & 255),
	// 			B: uint8((x + y) << 2 & 255),
	// 			A: 255,
	// 		})
	// 	}
	// }

	// f, err := os.Create("image.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// png.Encode(f, img)
	// f.Close()
}
