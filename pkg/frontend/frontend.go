package frontend

import (
	"log"

	tc "github.com/gdamore/tcell/v2"

	"github.com/lucidmachine/bitclock-go/pkg/backend"
)

func style(bit int) tc.Style {
	if bit > 0 {
		return tc.StyleDefault.Background(tc.ColorWhite)
	} else {
		return tc.StyleDefault.Background(tc.ColorBlack)
	}
}

func Init(screen tc.Screen) {
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(tc.StyleDefault.Background(tc.ColorBlack).Foreground(tc.ColorWhite))
	screen.Clear()
}

func Render(screen tc.Screen, bt backend.BitTime) {
	for x := 0; x < 6; x++ {
		for y := 0; y < 4; y++ {
			screen.SetContent(x, y, ' ', nil, style(bt[x][y]))
		}
	}

	screen.Show()
}
