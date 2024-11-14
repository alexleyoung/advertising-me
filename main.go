package main

import (
	"go-game/game"
	"go-game/game/screens"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func InitLogs() *os.File {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	return file
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	InitLogs()
	game.InitSaves()
	
	running := true
	for running {
		action := screens.MainMenu(screen)
		switch action.Type {
		case "PLAY":
			g := game.InitGame(action.Data)
			screens.Hub(screen, g)
		case "STATS":
			screens.Stats(screen)
		case "EXIT":
			running = false
		}
	}
}