package screens

import (
	"go-game/game"

	"github.com/gdamore/tcell/v2"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Stats(screen tcell.Screen) {
	// fetch and tabulate scores
	t := table.NewWriter()
	t.AppendHeader(table.Row{"PLAYER", "SCORE", "NEAR MISSES", "TIMESTAMP"})
	scores := game.GetHighScores()
	for _, score := range scores {
		t.AppendRow(table.Row{score.Player, score.Score, score.NearMisses, score.Timestamp})
	}

	screen.Clear()

	game.DrawTable(screen, 0, 0, t)
	
	screen.Show()

	for {

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				return
			}
		}	
	}
}