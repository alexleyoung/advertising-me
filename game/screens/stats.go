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
	scores := game.GetHighScores()
	for _, score := range scores {
		t.AppendRow(table.Row{score.Player, score.Score, score.NearMisses})
	}
	stats := true
	for stats {
		screen.Clear()

		game.DrawTable(screen, 0, 0, t)
		// game.DrawString(screen, 0, 0, "PLAYER | SCORE | NEAR MISSES | TIMESTAMP")
		// for i, score := range scores {
		// 	game.DrawString(screen, 0, 1+i, score.Player + 
		// 		" | " + strconv.Itoa(score.Score) + 
		// 		" | " + strconv.Itoa(score.NearMisses) + 
		// 		" | " + time.Unix(int64(score.Timestamp), 0).Format("Jan 2 2006 03:04 PM"))
		// } 

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