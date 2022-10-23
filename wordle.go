package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	title  string = "Wordle"
	width  int    = 435
	height int    = 600
	rows   int    = 6
	cols   int    = 5
)

var (
	fontSize        int = 6
	mplusNormalFont font.Face

	bkg       = color.White
	lightgrey = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	grey      = color.RGBA{0x77, 0x7c, 0x7e, 0xff}
	yellow    = color.RGBA{0xcd, 0xb3, 0x5d, 0xff}
	green     = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	fontColor = color.Black

	edge     = false
	alphabet = "qwertyuioplkjhgfdsazxcvbnm"
	grid     [cols * rows]string
	check    [cols * rows]int
	dict     []string
	loc      int = 0
	won          = false
	answer   string
)

type Game struct {
	runes []rune
}

func (g Game) Update() error {
	return nil
}

func RepeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	d := inpututil.KeyPressDuration(key)
	fmt.Println(d)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)&interval == 0 {
		return true
	}
	return true
}

func (g Game) Draw(screen *ebiten.Image) {
	for w := 0; w < cols; w++ {
		for h := 0; h < rows; h++ {
			rect := ebiten.NewImage(75, 75)
			rect.Fill(lightgrey)
			fontColor = color.Black
			if check[w+(h*cols)] != 0 {
				if check[w+(h*cols)] == 1 {
					rect.Fill(green)
				}
				if check[w+(h*cols)] == 2 {
					rect.Fill(yellow)
				}
				if check[w+(h*cols)] == 2 {
					rect.Fill(grey)
				}
				fontColor = color.White
			}
		}
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	g := &Game{}
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	file, err := os.ReadFile("wordle.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		dict = strings.Split(string(file), "\n")
	}
	rand.Seed(time.Now().UnixNano())
	answer = dict[rand.Intn(len(dict))]

	fmt.Println(answer)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
