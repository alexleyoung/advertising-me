package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Panicln(err)
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}
	// game init
	player := NewSprite('@', 10, 10)

	running := true
	for running {
		// update logic
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				running = false
			}
			switch ev.Rune() {
			case 'w':
				player.Y--
			case 's':
				player.Y++
			case 'a':
				player.X--
			case 'd':
				player.X++
			}
		}

		// draw logic
		screen.Clear()

		player.Draw(screen)

		screen.Show()
	}
}
