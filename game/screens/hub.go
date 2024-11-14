package screens

import (
	"go-game/game"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func Hub(screen tcell.Screen, g *game.Game) *Action {
	fps := 0
	frameCount := 0
	lastFPSUpdate := time.Now()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	PLAY_PORTAL_X := 15
	PLAY_PORTAL_Y := 25

	// game loop
	for {
		// draw logic
		screen.Clear()

		// draw player
		g.Player.Sprite.Draw(screen)
		// draw fps
		game.DrawString(screen, 0, 0, strconv.Itoa(fps))
		// draw game portal
		screen.SetContent(PLAY_PORTAL_X, PLAY_PORTAL_Y, '=', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		screen.SetContent(PLAY_PORTAL_X-1, PLAY_PORTAL_Y, '<', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		screen.SetContent(PLAY_PORTAL_X+1, PLAY_PORTAL_Y, '>', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))

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
		if g.Player.Sprite.Y == PLAY_PORTAL_Y && g.Player.Sprite.X == PLAY_PORTAL_X {
			Game(screen, g)
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
