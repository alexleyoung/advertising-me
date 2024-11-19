package screens

import (
	"go-game/game"

	"github.com/gdamore/tcell/v2"
)

type Action struct {
	Type string
	Data string
}

func MainMenu(screen tcell.Screen) *Action {
	mainMenu := true
	players := game.GetPlayers()
	selected := 0
	playerName := ""
	lastInput := tcell.KeyRune
	art := game.ImgToAscii("assets/headshot.png", 55, 35)
	art2 := game.ImgToAscii("assets/now/alexcolin.jpg", 80, 35)
	for mainMenu {
		screen.Clear()

		game.DrawString(screen, 0, 0, art)
		game.DrawString(screen, 90, 0, art2)
		// draw menu UI
		game.DrawString(screen, 60, 20, "Welcome to Advertising Alex!")
		game.DrawString(screen, 67, 22, "Who is playing?")
		for i, player := range players {
			if i == selected {
				game.DrawColorString(screen, 69, 24+i, player, tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			} else {
				game.DrawString(screen, 69, 24+i, player)
			}
		}
		game.DrawString(screen, 57, 24+len(players), "new player: ")
		if selected == len(players) {
			if len(playerName) == 0 {
				game.DrawColorString(screen, 69, 24+len(players), "_", tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			} else {
				game.DrawColorString(screen, 69, 24+len(players), playerName, tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			}
		}
		if lastInput == 127 && selected != len(players) {
			game.DrawColorString(screen, 55, 30, "Press delete again to remove selected player", tcell.StyleDefault.Foreground(tcell.ColorRed))
		}

		screen.Show()

		// handle inputs
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				mainMenu = false
				return &Action{
					Type: "EXIT",
					Data: "",
				}
			case tcell.KeyTAB:
				return &Action{
					Type: "STATS",
					Data: "",
				}
			case tcell.KeyDown:
				if selected < len(players) {
					selected++
					lastInput = 0
				}
			case tcell.KeyUp:
				if selected > 0 {
					selected--
					lastInput = 0
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				backspaced := false
				if selected != len(players) {
					if lastInput == 127 {
						game.RemovePlayer(players[selected])
						players = append(players[:selected], players[selected+1:]...)
						if selected >= len(players) {
							selected = len(players) - 1
						}
						backspaced = true
					}
				}
				if len(playerName) > 0 {
					playerName = playerName[:len(playerName)-1]
				}
				lastInput = ev.Key()
				if backspaced {
					lastInput = 0
				}
			case tcell.KeyRune:
				lastInput = 0
				if selected < len(players) || ev.Rune() == ' ' {
					break
				}
				playerName += string(ev.Rune())
			case tcell.KeyEnter:
				if selected < len(players) {
					game.CreatePlayer(players[selected])
					mainMenu = false
					return &Action{
						Type: "PLAY",
						Data: players[selected],
					}
				} else {
					if len(playerName) > 0 {
						game.CreatePlayer(playerName)
						mainMenu = false
						return &Action{
							Type: "PLAY",
							Data: playerName,
						}
					}
				}
			}
		}
	}
	return &Action{
		Type: "EXIT",
		Data: "",
	}
}
