package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func reverseSet(set []rune) []rune {
	reversed := []rune{}
	for i := len(set) - 1; i >= 0; i-- {
		reversed = append(reversed, set[i])
	}
	return reversed
}

func squash(value float64) float64 {
	return math.Pow(value, 3) / 255 / 255
}

func grayToChar(c color.Gray, characterSet []rune) string {
	levels := float64(len(characterSet))
	level := squash(float64(c.Y)) / 255.0 * levels
	if level == levels {
		level--
	}
	return string(characterSet[int(level)])
}

func printImage(img image.Image, characterSet []rune, scale int) {
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y += scale {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x += scale / 2 {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			fmt.Print(grayToChar(c, characterSet))
		}
		fmt.Print("\n")
	}
}

func main() {
	var hFlag = flag.Bool("h", false, "")
	var iFlag = flag.Bool("i", false, "")
	var fFlag = flag.String("f", "", "")
	var sFlag = flag.Int("s", 20, "")
	var cFlag = flag.String("c", "standard", "")
	flag.Parse()
	_ = *iFlag
	_ = sFlag
	_ = fFlag
	_ = cFlag

	sets := map[string][]rune{
		"standard": []rune(" .'" + "`" + "^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ1OZmwqpdbkhao*#MW&8%B@$"),
		"detailed": []rune(" `.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@"),
		"simple":   []rune(" .:-=+*#%@"),
		"squares":  []rune(" ░▒▓█"),
	}

	// afFlag := "~/Pictures/other/1684441894831350.jpg"

	if *hFlag {
		fmt.Println("idk...")
		os.Exit(0)
	}

	characterSet := sets[*cFlag]

	if *iFlag {
		characterSet = reverseSet(characterSet)
	}

	reader, err := os.Open(*fFlag)

	if err != nil {
		log.Fatal(err)
	}
	// _, imageType, err := image.Decode(reader)
	//
	// if err != nil {
	// 	fmt.Println(filePath)
	// 	log.Fatal(err)
	// }
	//
	img, err := jpeg.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	printImage(img, characterSet, *sFlag)
}
