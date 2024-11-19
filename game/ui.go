package game

import (
	"github.com/alexleyoung/quickscii"
	"github.com/gdamore/tcell/v2"
	"github.com/jedib0t/go-pretty/v6/table"
)

func DrawString(screen tcell.Screen, x, y int, s string) {
	i := x
	for _, r := range s {
		if r == '\n' {
			y++
			i = x
			continue
		}
		screen.SetContent(i, y, r, nil, tcell.StyleDefault)
		i++
	}
}

func DrawColorString(screen tcell.Screen, x, y int, s string, color tcell.Style) {
	for _, r := range s {
		if r == '\n' {
			y++
			x = 0
			continue
		}
		screen.SetContent(x, y, r, nil, color)
		x++
	}
}

func DrawRect(screen tcell.Screen, x, y, w, h int, style tcell.Style) {
	for j := y; j < y+h; j++ {
		screen.SetContent(x, j, '|', nil, style)
		screen.SetContent(x+w, j, '|', nil, style)
	}
	for i := x; i < x+w+1; i++ {
		screen.SetContent(i, y, '-', nil, style)
		screen.SetContent(i, y+h, '-', nil, style)
	}
}

func DrawTable(screen tcell.Screen, x, y int, t table.Writer) {
	t.SetStyle(table.StyleLight)

	content := t.Render()
	DrawString(screen, x, y, content)
}

func ImgToAscii(path string, x, y int) string {
	return quickscii.Convert(path, x, y)
}
