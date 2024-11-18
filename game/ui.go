package game

import (
	"github.com/alexleyoung/quickscii/quickscii"
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

func DrawTable(screen tcell.Screen, x, y int, t table.Writer) {
	t.SetStyle(table.StyleLight)

	content := t.Render()
	DrawString(screen, x, y, content)
}

func ImgToAscii(path string, x, y int) string {
    return quickscii.Convert(path, x, y)
}