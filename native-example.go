package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/paint"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case paint.Event:
				// call opengl here!
				_ = e
			}
		}
	})
}
