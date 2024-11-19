package screens

import (
	"go-game/game"

	"github.com/gdamore/tcell/v2"
)

func Slides(screen tcell.Screen, imgs []*Image) {
	for _, img := range imgs {
		// first picture
		screen.Clear()

		// scuffed logic for loading pictures side by side
		if img.Path == "assets/childhood/steam.png" {
			side, err := game.ImgToAscii("assets/childhood/league.png", 80, 40, img.Charset)
			if err != nil {
				panic(err)
			}
			game.DrawString(screen, 81, 10, side)
		}
		if img.Path == "assets/childhood/band.JPG" {
			side, err := game.ImgToAscii("assets/childhood/eigth.JPG", 80, 50, img.Charset)
			if err != nil {
				panic(err)
			}
			game.DrawString(screen, 81, 0, side)
		}

		// draw ascii picture
		img, err := game.ImgToAscii(img.Path, img.Width, img.Height, img.Charset)
		if err != nil {
			panic(err)
		}
		game.DrawString(screen, 0, 0, img)
		screen.Show()
		run := true
		for run {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEnter:
					run = false
				}
			}
		}
	}
}
