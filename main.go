package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

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
	level := 1
	coins := generateCoins(level)	
	score := 0
	fps := 0
	frameCount := 0
	lastFPSUpdate := time.Now()

	running := true
	for running {
		// update logic
		playerMoved := false
		// start := time.Now()

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
					playerMoved = true
				}
			case 'j', 's':
				if player.Y < 42 {
					player.Y++
					playerMoved = true
				}
			case 'h', 'a':
				if player.X > 0 {
					player.X--
					playerMoved = true
				}
			case 'l', 'd':
				if player.X < 130 {
					player.X++
					playerMoved = true
				}
			}
		}

		if playerMoved {
			for i, coin := range coins {
				if coin.X == player.X && coin.Y == player.Y {
					coins[i] = coins[len(coins)-1]
					coins = coins[:len(coins)-1]
					score++
					if len(coins) == 0 {
						level++
						coins = generateCoins(level)
						score = 0
					}
					break
				}
			}
		}

		// draw logic
		screen.Clear()

		player.Draw(screen)
		for _, coin := range coins {
			coin.Draw(screen)
		}
		drawString(screen, 0, 0, fmt.Sprintf("Level: %d", level))
		drawString(screen, 0, 1, fmt.Sprintf("Coins: %d/%d", score, level+2))
		drawString(screen, 0, 2, fmt.Sprintf("FPS: %d", fps))

		screen.Show()
		

		frameCount++
		if time.Since(lastFPSUpdate) >= time.Second {
			fps = frameCount
			frameCount = 0
			lastFPSUpdate = time.Now()
		}
	}
}
