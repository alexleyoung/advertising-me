package game

import (
	"github.com/alexleyoung/quickscii/quickscii"
	"github.com/gdamore/tcell/v2"
	"github.com/jedib0t/go-pretty/v6/table"
)

func DrawString(screen tcell.Screen, x, y int, s string) {
	for i, r := range s {
		screen.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}
func DrawColorString(screen tcell.Screen, x, y int, s string, color tcell.Style) {
	for i, r := range s {
		screen.SetContent(x+i, y, r, nil, color)
	}
}

func DrawTable(screen tcell.Screen, x, y int, t table.Writer) {
	t.SetStyle(table.StyleLight)
	t.Style().Size.WidthMax = 20

	content := t.Render()
	
	row := x
	col := y
	for _, char := range content {
		if char == '\n' {
			col += 1
			row = x
		}
		screen.SetContent(row, col, char, nil, tcell.StyleDefault)
	}
}

func ImgToAscii(path string, x, y int) string {
	mat, err := quickscii.PreProcess(path, x, y)
	if err != nil {
		return ""
	}
	return quickscii.Convert(mat)
}