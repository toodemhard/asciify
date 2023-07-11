package asciify

import (
	"flag"
	"fmt"
	"image"
	"log"
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

func Start() {
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

	printImage(img, characterSet, scaleToSteps(*sFlag, img.Bounds().Max.Y))
}
