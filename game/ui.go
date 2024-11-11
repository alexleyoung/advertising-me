package game

import (
	"github.com/gdamore/tcell/v2"
)

func DrawString(screen tcell.Screen, x, y int, s string) {
	for i, r := range s {
		screen.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}