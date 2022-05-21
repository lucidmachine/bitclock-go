package main

import (
	"log"
	"os"
	"time"

	tc "github.com/gdamore/tcell/v2"

	"github.com/lucidmachine/bitclock-go/pkg/backend"
	"github.com/lucidmachine/bitclock-go/pkg/frontend"
)

func handleEvent(screen tc.Screen, ev tc.Event) {
	switch ev := ev.(type) {

	case *tc.EventResize:
		screen.Sync()

	case *tc.EventKey:
		if ev.Key() == tc.KeyEscape || ev.Key() == tc.KeyCtrlC {
			screen.Fini()
			os.Exit(0)
		}
	}
}

func main() {
	screen, err := tc.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	frontend.Init(screen)

	for {
		// Build and display time
		now := time.Now()
		bt := backend.NewBitTime(now)
		frontend.Render(screen, bt)

		// Handle events
		if screen.HasPendingEvent() {
			ev := screen.PollEvent()
			handleEvent(screen, ev)
		}

		// Wait before next render
		time.Sleep(1 * time.Second)
	}
}
