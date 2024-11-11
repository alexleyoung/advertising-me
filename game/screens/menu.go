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
	playerName := ""
	for mainMenu {
		screen.Clear()

		game.DrawString(screen, 60, 20, "Welcome to Advertising Alex!")
		game.DrawString(screen, 67, 22, "Who is playing?")
		game.DrawString(screen, 69, 24, playerName)

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
			case tcell.KeyEnter:
				if len(playerName) > 0 {
					game.CreatePlayer(playerName)
					mainMenu = false
					return &Action{
						Type: "PLAY",
						Data: playerName,
					}
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(playerName) > 0 {
					playerName = playerName[:len(playerName)-1]
				}
			case tcell.KeyTAB:
				// go to stats screen
			case tcell.KeyRune:
				playerName += string(ev.Rune())
			}
		}	
	}
	return &Action{
		Type: "EXIT",
		Data: "",
	}
}