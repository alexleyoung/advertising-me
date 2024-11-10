package main

import (
	"fmt"
	"log"

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

	running := true
	for running {
		// update logic
		playerMoved := false

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
				playerMoved = true
			case 's':
				player.Y++
				playerMoved = true
			case 'a':
				player.X--
				playerMoved = true
			case 'd':
				player.X++
				playerMoved = true
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

		screen.Show()
	}
}
