package screens

import (
	"go-game/game"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func Shop(screen tcell.Screen, g *game.Game, coins int) {
	SHOP_TEXT := `
	███████  ██   ██  ███████   ██████ 
	██       ██   ██  ██   ██   ██  ██
	███████  ███████  ██   ██   ██████
	     ██  ██   ██  ██   ██   ██
	███████  ██   ██  ███████   ██
	----------------------------------
	`
	MAP_HEIGHT := 30
	MAP_WIDTH := 78

	for {	
		// draw logic
		screen.Clear()

		// draw ui
		game.DrawString(screen, 0, 0, "Coins: " + strconv.Itoa(coins))

		g.Player.Sprite.Draw(screen)

		// draw map
		game.DrawString(screen, 63, 0, SHOP_TEXT)
		game.DrawRect(screen, 40, 10, MAP_WIDTH, MAP_HEIGHT, tcell.StyleDefault)
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
					if g.Player.Sprite.Y > 10 {
						g.Player.Sprite.Y--
					}
				case 'j', 's':
					if g.Player.Sprite.Y < 10 + MAP_HEIGHT {
						g.Player.Sprite.Y++
					}
				case 'h', 'a':
					if g.Player.Sprite.X > 40 {
						g.Player.Sprite.X--
					}
				case 'l', 'd':
					if g.Player.Sprite.X < 40+MAP_WIDTH {
						g.Player.Sprite.X++
					}
				}
			}
		}	
	}
}