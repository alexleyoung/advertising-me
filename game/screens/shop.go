package screens

import (
	"go-game/game"

	"github.com/gdamore/tcell/v2"
)

func Shop(screen tcell.Screen, g *game.Game) {
	for {	
		// draw logic
		screen.Clear()

		g.Player.Sprite.Draw(screen)

		screen.Show()

		// movement
		if screen.HasPendingEvent() {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return 
				}
				switch ev.Rune() {
				case 'k', 'w':
					if g.Player.Sprite.Y > 0 {
						g.Player.Sprite.Y--
					}
				case 'j', 's':
					if g.Player.Sprite.Y < 42 {
						g.Player.Sprite.Y++
					}
				case 'h', 'a':
					if g.Player.Sprite.X > 0 {
						g.Player.Sprite.X--
					}
				case 'l', 'd':
					if g.Player.Sprite.X < 130 {
						g.Player.Sprite.X++
					}
				}
			}
		}	
	}
}