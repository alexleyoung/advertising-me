package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)
type Projectile struct {
	Sprite *Sprite
	SpeedX, SpeedY int
}

func NewProjectile(x, y, sx, sy int) *Projectile {
	return &Projectile{
		Sprite: NewSprite('*', x, y),
		SpeedX: sx,
		SpeedY: sy,
	}
}

func (p *Projectile) Update() {
	p.Sprite.X += p.SpeedX
	p.Sprite.Y += p.SpeedY
}

func drawString(screen tcell.Screen, x, y int, s string) {
	for i, r := range s {
		screen.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}

func generateCoins(level int) []*Sprite {
	coins := make([]*Sprite, level + 2)

	for i := range level + 2 {
		coins[i] = NewSprite('o', rand.Intn(120) + 5, rand.Intn(32) + 5)
	}

	return coins
}

func generateProjectile() *Projectile {
	spawn := rand.Intn(4)
	var x, y, sx, sy int
	switch spawn {
	case 0:
		x = rand.Intn(120) + 5
		y = rand.Intn(5) - 5
		sx = 0
		sy = 1
	case 1:
		x = rand.Intn(5) + 140
		y = rand.Intn(32) + 5
		sx = -1
		sy =0
	case 2:
		x = rand.Intn(120) + 5
		y = rand.Intn(5) + 45
		sx = 0
		sy = -1
	case 3:
		x = rand.Intn(5) - 5 
		y = rand.Intn(32) + 5
		sx = 1
		sy = 0
	}
	return NewProjectile(x, y, sx, sy)
}

func generateProjectiles(level int) []*Projectile {
	projectiles := make([]*Projectile, level * 4)

	for i := range level * 4 {
		projectiles[i] = generateProjectile()	
	}

	return projectiles
}

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
	player := NewSprite('@', 70, 20)
	level := 1
	coins := generateCoins(level)
	projectiles := generateProjectiles(level)
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
			if playerMoved {
				for i, coin := range coins {
					if coin.X == player.X && coin.Y == player.Y {
						coins[i] = coins[len(coins)-1]
						coins = coins[:len(coins)-1]
						score++
						if len(coins) == 0 {
							level++
							coins = generateCoins(level)
							projectiles = generateProjectiles(level)
							score = 0
						}
						break
					}
				}
			}
			
			for i := len(projectiles) - 1; i >= 0; i-- {
				projectile := projectiles[i]
				projectile.Update()
				if projectile.Sprite.X < -5 || projectile.Sprite.X > 150 || projectile.Sprite.Y < -5 || projectile.Sprite.Y > 50 {
					projectiles[i] = generateProjectile()
				}
				if projectile.Sprite.Y == player.Y && projectile.Sprite.X == player.X {
					alive = false
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
	
			drawString(screen, 0, 0, fmt.Sprintf("Level: %d", level))
			drawString(screen, 0, 1, fmt.Sprintf("Coins: %d/%d", score, level+2))
			drawString(screen, 0, 2, fmt.Sprintf("FPS: %d", fps))
	
			screen.Show()
		} else {
			drawString(screen, 70, 20, "GAME OVER")
			drawString(screen, 63, 22, "Press any key to restart")
			screen.Show()
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					running = false
				case tcell.KeyRune:
					// reinitialize game
					player = NewSprite('@', 70, 20)
					level = 1
					coins = generateCoins(level)
					projectiles = generateProjectiles(level)
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