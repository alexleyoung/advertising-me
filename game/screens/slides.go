package screens

import (
	"go-game/game"

	"github.com/gdamore/tcell/v2"
)

func Slides(screen tcell.Screen, imgs []*Image) {
	for _, img := range imgs {
		// first picture
		screen.Clear()
		if img.Path == "assets/childhood/steam.png" {
			game.DrawString(screen, 81, 10, game.ImgToAscii("assets/childhood/league.png", 80, 40))
		}
		if img.Path == "assets/childhood/band.JPG" {
			game.DrawString(screen, 81, 0, game.ImgToAscii("assets/childhood/eigth.JPG", 80, 50))
		}
		img := game.ImgToAscii(img.Path, img.Width, img.Height)
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
