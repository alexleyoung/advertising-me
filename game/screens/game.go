package screens

import (
	"fmt"
	"math"
	"time"

	"go-game/game"

	"github.com/gdamore/tcell/v2"
)

func Game(screen tcell.Screen, g *game.Game) {
	coinColor := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	projectileColor := tcell.StyleDefault.Foreground(tcell.ColorRed)

	// fps counter initialization
	fps := 0
	frameCount := 0
	lastFPSUpdate := time.Now()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	for {
		if g.Alive {
			// Player movement
			playerMoved := false
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
						if g.Player.Sprite.Y > game.TOP_BORDER_Y+1 {
							g.Player.Sprite.Y--
						}
						playerMoved = true
					case 'j', 's':
						if g.Player.Sprite.Y < game.TOP_BORDER_Y+game.MAP_HEIGHT-1 {
							g.Player.Sprite.Y++
						}
						playerMoved = true
					case 'h', 'a':
						if g.Player.Sprite.X > game.LEFT_BORDER_X+1 {
							g.Player.Sprite.X--
						}
						playerMoved = true
					case 'l', 'd':
						if g.Player.Sprite.X < game.LEFT_BORDER_X+game.MAP_WIDTH-1 {
							g.Player.Sprite.X++
						}
						playerMoved = true
					}
				}
			}

			// coin collision check
			if playerMoved {
				for i, coin := range g.Coins {
					if coin.X == g.Player.Sprite.X && coin.Y == g.Player.Sprite.Y {
						g.Coins[i] = g.Coins[len(g.Coins)-1]
						g.Coins = g.Coins[:len(g.Coins)-1]
						g.CoinCount++
						g.Player.Score++
						g.Player.Coins++
						if len(g.Coins) == 0 {
							g.Level++
							g.Coins = game.GenerateCoins(g.Level, coinColor)
							g.Projectiles = game.GenerateProjectiles(g.Level, projectileColor)
							g.CoinCount = 0
						}
						break
					}
				}
			}

			// projectile collision check
			for i := len(g.Projectiles) - 1; i >= 0; i-- {
				projectile := g.Projectiles[i]
				projectile.Update()

				// respawn out of bounds projectiles
				if projectile.Sprite.X < -5 || projectile.Sprite.X > 150 || projectile.Sprite.Y < -5 || projectile.Sprite.Y > 50 {
					g.Projectiles[i] = game.GenerateProjectile(projectileColor)
				}
				if projectile.Sprite.Y == g.Player.Sprite.Y && projectile.Sprite.X == g.Player.Sprite.X {
					g.Alive = false
				}

				// check for near miss
				if math.Abs(float64(projectile.Sprite.X-g.Player.Sprite.X)) == 1 && math.Abs(float64(projectile.Sprite.Y-g.Player.Sprite.Y)) == 1 {
					g.Player.Score++
					g.Player.NearMisses++
				}
			}

			// draw logic
			screen.Clear()

			// draw map border
			game.DrawRect(screen, game.LEFT_BORDER_X, game.TOP_BORDER_Y, game.MAP_WIDTH, game.MAP_HEIGHT, tcell.StyleDefault)

			// draw player, coins, projectilees
			g.Player.Sprite.Draw(screen)
			for _, coin := range g.Coins {
				coin.Draw(screen)
			}
			for _, projectile := range g.Projectiles {
				projectile.Sprite.Draw(screen)
			}

			// draw ui
			game.DrawString(screen, 0, 0, g.Player.Name)
			game.DrawString(screen, 0, 1, fmt.Sprintf("Score: %d", g.Player.Score))
			game.DrawString(screen, 0, 2, fmt.Sprintf("Level: %d", g.Level))
			game.DrawString(screen, 0, 3, fmt.Sprintf("Coins 🪙: %d/%d", g.CoinCount, g.Level+2))
			game.DrawString(screen, 147, 0, fmt.Sprintf("FPS: %d", fps))

			screen.Show()

			// fps counter logic
			frameCount++
			if time.Since(lastFPSUpdate) >= time.Second {
				fps = frameCount
				frameCount = 0
				lastFPSUpdate = time.Now()
			}

			<-ticker.C
		} else {
			// player death logic
			// save score
			game.SavePlayerData(g.Player.Name, g.Player.Score, g.Player.NearMisses, time.Now().Unix())
			game.AddItem(g.Player.Name, "coin", g.Player.Coins)

			// draw game over
			game.DrawString(screen, 70, 20, "GAME OVER")
			game.DrawString(screen, 50, 22, "Press any key to restart, or ESC to go back to the hub")
			screen.Show()

			// reinitialize game
			g = game.InitGame(g.Player.Name)

			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					g.Alive = true
					return
				case tcell.KeyRune:
					// reinitialize game
					g.Alive = true
				}
			}
		}
	}
}
