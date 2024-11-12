package screens

import (
	"go-game/game"

	"github.com/gdamore/tcell/v2"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Stats(screen tcell.Screen) {
	// fetch and tabulate scores
	t := table.NewWriter()
	t.AppendHeader(table.Row{"PLAYER", "SCORE", "NEAR MISSES"})
	for _, score := range game.GetHighScores() {
		t.AppendRow(table.Row{score.Player, score.Score, score.NearMisses})
	}
	t.SetStyle(table.StyleLight)
	t.Style().Size.WidthMax = 50
	t.Style().Size.WidthMin = 50
	
	stats := true
	for stats {
		screen.Clear()

		screen.Show()

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				stats = false
			}
		}	
	}
}