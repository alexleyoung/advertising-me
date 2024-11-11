package main

import (
	"fmt"
	"log"
	"math"
	"time"

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

	playerColor := tcell.StyleDefault.Foreground(tcell.ColorTeal)
	coinColor := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	projectileColor := tcell.StyleDefault.Foreground(tcell.ColorRed)

	// game init
	player := NewSprite('@', 70, 20, playerColor)
	level := 1
	coins := GenerateCoins(level, coinColor)
	projectiles := GenerateProjectiles(level, projectileColor)
	coinCount := 0
	score := 0
	running := true
	alive := true
	
	// manage fps
	fps := 0
	frameCount := 0
	lastFPSUpdate := time.Now()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	for running {
		// update logic
		playerMoved := false
		
		if screen.HasPendingEvent() {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					running = false
				}
				switch ev.Rune() {
				case 'k', 'w':
					if player.Y > 0 {
						player.Y--
					}
					playerMoved = true
				case 'j', 's':
					if player.Y < 42 {
						player.Y++
					}
					playerMoved = true
				case 'h', 'a':
					if player.X > 0 {
						player.X--
					}
					playerMoved = true
				case 'l', 'd':
					if player.X < 130 {
						player.X++
					}
					playerMoved = true
				}
			}
		}	
		if alive {
			// coin collision check
			if playerMoved {
				for i, coin := range coins {
					if coin.X == player.X && coin.Y == player.Y {
						coins[i] = coins[len(coins)-1]
						coins = coins[:len(coins)-1]
						coinCount++
						score++
						if len(coins) == 0 {
							level++
							coins = GenerateCoins(level, coinColor)
							projectiles = GenerateProjectiles(level, projectileColor)
							coinCount = 0
						}
						break
					}
				}
			}
			
			// projectile collision check
			for i := len(projectiles) - 1; i >= 0; i-- {
				projectile := projectiles[i]
				projectile.Update()

				// respawn out of bounds projectiles
				if projectile.Sprite.X < -5 || projectile.Sprite.X > 150 || projectile.Sprite.Y < -5 || projectile.Sprite.Y > 50 {
					projectiles[i] = GenerateProjectile(projectileColor)
				}
				if projectile.Sprite.Y == player.Y && projectile.Sprite.X == player.X {
					alive = false
				}

				// chcek for near miss
				if math.Abs(float64(projectile.Sprite.X - player.X)) == 1 && math.Abs(float64(projectile.Sprite.Y - player.Y)) == 1 {
					score++
				}
			}
	
			// draw logic
			screen.Clear()
	
			player.Draw(screen)
			for _, coin := range coins {
				coin.Draw(screen)
			}
			for _, projectile := range projectiles {
				projectile.Sprite.Draw(screen)
			}
			DrawString(screen, 0, 0, fmt.Sprintf("Score: %d", score))
			DrawString(screen, 0, 1, fmt.Sprintf("Level: %d", level))
			DrawString(screen, 0, 2, fmt.Sprintf("Coins: %d/%d", coinCount, level+2))
			DrawString(screen, 0, 3, fmt.Sprintf("FPS: %d", fps))
	
			screen.Show()
		} else {
			DrawString(screen, 70, 20, "GAME OVER")
			DrawString(screen, 63, 22, "Press any key to restart")
			screen.Show()
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					running = false
				case tcell.KeyRune:
					// reinitialize game
					player = NewSprite('@', 70, 20, playerColor)
					level = 1
					coins = GenerateCoins(level, coinColor)
					projectiles = GenerateProjectiles(level, projectileColor)
					coinCount = 0
					score = 0
					alive = true
					running = true
				}
			}
		}
		

		frameCount++
		if time.Since(lastFPSUpdate) >= time.Second {
			fps = frameCount
			frameCount = 0
			lastFPSUpdate = time.Now()
		}

		<-ticker.C
	}
}