package screens

import (
	"go-game/game"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func Hub(screen tcell.Screen, g *game.Game) {
	running := true

	fps := 0
	frameCount := 0
	lastFPSUpdate := time.Now()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	// game loop
	for running {
		// draw logic
		screen.Clear()

		// draw player
		g.Player.Sprite.Draw(screen)
		// draw fps
		game.DrawString(screen, 0, 0, strconv.Itoa(fps))
		// draw game portal
		screen.SetContent(15,25, '=', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		screen.SetContent(14,25, '<', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		screen.SetContent(16,25, '>', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))

		screen.Show()

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
					if g.Player.Sprite.Y > 0 {
						g.Player.Sprite.Y--
					}
					playerMoved = true
				case 'j', 's':
					if g.Player.Sprite.Y < 42 {
						g.Player.Sprite.Y++
					}
					playerMoved = true
				case 'h', 'a':
					if g.Player.Sprite.X > 0 {
						g.Player.Sprite.X--
					}
					playerMoved = true
				case 'l', 'd':
					if g.Player.Sprite.X < 130 {
						g.Player.Sprite.X++
					}
					playerMoved = true
				}
			}
		}	

		if playerMoved {
			// check collisions with portals
			if g.Player.Sprite.Y == 25 && g.Player.Sprite.X == 15 {
				Game(screen, g)
			}
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