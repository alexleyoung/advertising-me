package screens

import (
	"go-game/game"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func Hub(screen tcell.Screen, g *game.Game) *Action {
	// fps counter initialization
	fps := 0
	frameCount := 0
	lastFPSUpdate := time.Now()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	// portal coordinates
	PLAY_PORTAL_X := 35
	PLAY_PORTAL_Y := 25

	SHOP_PORTAL_X := 70
	SHOP_PORTAL_Y := 10

	// fetch player's coins
	coins := game.GetCoins(g.Player.Name)

	// game loop
	for {
		// draw logic
		screen.Clear()

		// draw player
		g.Player.Sprite.Draw(screen)
		// draw ui
		game.DrawString(screen, 0, 0, "Coins 🪙: "+strconv.Itoa(coins))
		game.DrawString(screen, 147, 0, strconv.Itoa(fps))
		// draw game portal
		screen.SetContent(PLAY_PORTAL_X, PLAY_PORTAL_Y, '🚪', nil, tcell.StyleDefault.Foreground(tcell.ColorPaleTurquoise))
		// draw shop portal
		screen.SetContent(SHOP_PORTAL_X, SHOP_PORTAL_Y, '🏪', nil, tcell.StyleDefault.Foreground(tcell.ColorFireBrick))

		screen.Show()

		// update logic
		if screen.HasPendingEvent() {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return &Action{
						Type: "EXIT",
						Data: "",
					}
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

		// check collisions with portals
		if g.Player.Sprite.X == PLAY_PORTAL_X && g.Player.Sprite.Y == PLAY_PORTAL_Y {
			g.Player.Sprite.X = 75
			g.Player.Sprite.Y = 25
			Game(screen, g)
			g = game.InitGame(g.Player.Name)
			coins = game.GetCoins(g.Player.Name)
		}
		if g.Player.Sprite.X == SHOP_PORTAL_X && g.Player.Sprite.Y == SHOP_PORTAL_Y {
			g.Player.Sprite.X = 75
			g.Player.Sprite.Y = 25
			Shop(screen, g, coins)
			g.Player.Sprite.X = SHOP_PORTAL_X
			g.Player.Sprite.Y = SHOP_PORTAL_Y + 1
			coins = game.GetCoins(g.Player.Name)
		}

		// fps counter logic
		frameCount++
		if time.Since(lastFPSUpdate) >= time.Second {
			fps = frameCount
			frameCount = 0
			lastFPSUpdate = time.Now()
		}

		<-ticker.C
	}
}
