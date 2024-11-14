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
	for mainMenu {
		screen.Clear()

		game.DrawString(screen, 60, 20, "Welcome to Advertising Alex!")
		game.DrawString(screen, 67, 22, "Who is playing?")
		for i, player := range players {
			if i == selected {
				game.DrawColorString(screen, 69, 24+i, player, tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			} else {
				game.DrawString(screen, 69, 24+i, player)
			}
		}
		game.DrawString(screen, 57, 24 + len(players), "new player: ")
		if selected == len(players) {
			if len(playerName) == 0 {
				game.DrawColorString(screen, 69, 24 + len(players), "_", tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			} else {
				game.DrawColorString(screen, 69, 24 + len(players), playerName, tcell.StyleDefault.Foreground(tcell.ColorOrangeRed))
			}
		}
		screen.Show()

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
				}
			case tcell.KeyUp:
				if selected > 0 {
					selected--
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(playerName) > 0 {
					playerName = playerName[:len(playerName)-1]
				}
			case tcell.KeyRune:
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