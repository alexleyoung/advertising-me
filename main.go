package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

func drawString(screen tcell.Screen, x, y int, s string) {
	for i, r := range s {
		screen.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
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
	coins := []*Sprite{
		NewSprite('o', 10, 50),
		NewSprite('o', 20, 32),
		NewSprite('o', 30, 28),
		NewSprite('o', 40, 15),
	}
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
			case 'w':
				if player.Y > 0 {
					player.Y--
					playerMoved = true
				}
			case 's':
				if player.Y < 100 {
					player.Y++
					playerMoved = true
				}
			case 'a':
				if player.X > 0 {
					player.X--
					playerMoved = true
				}
			case 'd':
				if player.X < 100 {
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
		drawString(screen, 0, 0, fmt.Sprintf("Score: %d", score))
		drawString(screen, 0, 1, fmt.Sprintf("FPS: %d", fps))

		screen.Show()

		frameCount++
		if time.Since(lastFPSUpdate) >= time.Second {
			fps = frameCount
			frameCount = 0
			lastFPSUpdate = time.Now()
		}
	}
}
