package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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

func scaleToSteps(lineHeight int, imageHeight int) int {
	return imageHeight / lineHeight
}

func grayToChar(c color.Gray, characterSet []rune) string {
	levels := float64(len(characterSet))
	level := squash(float64(c.Y)) / 255.0 * levels
	if level == levels {
		level--
	}
	return string(characterSet[int(level)])
}

// func areaColor(img image.Image, x int, y int, width int, height int) color.Gray {
// 	for y; y < y+height; y++ {
// 		for x; x < x+width; x++ {
//
// 		}
// 	}
// }

func printImage(img image.Image, characterSet []rune, scaleSteps int) {
	// stringImage := ""
	halfStep := scaleSteps / 2
	for y := img.Bounds().Min.Y; y+scaleSteps < img.Bounds().Max.Y; y += scaleSteps {
		for x := img.Bounds().Min.X; x+halfStep < img.Bounds().Max.X; x += halfStep {
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

	img, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(img.Bounds().Max.Y)
	printImage(img, characterSet, scaleToSteps(*sFlag, img.Bounds().Max.Y))
}
