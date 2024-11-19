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
	for mainMenu {
		screen.Clear()

		// draw menu UI
		game.DrawString(screen, 60, 20, "Welcome to Advertising Alex!")
		game.DrawString(screen, 66, 22, "Who is playing?")

		// draw player list
		for i, player := range players {
			if i == selected {
				game.DrawColorString(screen, 70, 24+i, player, tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			} else {
				game.DrawString(screen, 70, 24+i, player)
			}
		}

		// draw new player creation
		game.DrawString(screen, 57, 24+len(players), "new player: ")
		if selected == len(players) {
			if len(playerName) == 0 {
				game.DrawColorString(screen, 70, 24+len(players), "_", tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			} else {
				game.DrawColorString(screen, 70, 24+len(players), playerName, tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			}
		}

		// draw delete warning
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
				// remove player if second backspace press
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
				// remove character from new player name
				if len(playerName) > 0 {
					playerName = playerName[:len(playerName)-1]
				}
				lastInput = ev.Key()
				if backspaced {
					lastInput = 0
				}
			case tcell.KeyRune:
				lastInput = 0
				// disable space for player names
				if selected < len(players) || ev.Rune() == ' ' {
					break
				}
				// append character to new player name
				playerName += string(ev.Rune())
			case tcell.KeyEnter:
				// select player
				if selected < len(players) {
					game.CreatePlayer(players[selected])
					mainMenu = false
					return &Action{
						Type: "PLAY",
						Data: players[selected],
					}
				} else { // create and select new player
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
