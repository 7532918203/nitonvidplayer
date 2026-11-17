package main

import (
	"fmt"
	"fyneVideoPlayer/parser"
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: animeplayer <library folder>")
		os.Exit(2)
	}
	library, _ := parser.ScanLibrary(os.Args[1])

	for _, v := range library.Shows {
		fmt.Println(v)
	}

	a := app.New()
	w := a.NewWindow("Update Time")

	message := widget.NewLabel("Welcome")
	button := widget.NewButton("Update", func() {
		formatted := time.Now().Format("Time: 03:04:05")
		message.SetText(formatted)
	})

	w.SetContent(container.NewVBox(message, button))
	w.ShowAndRun()
}
