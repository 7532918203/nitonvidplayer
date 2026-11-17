package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type showCell struct {
	widget.BaseWidget
	poster  *canvas.Rectangle
	initial *canvas.Text
	title   *widget.Label
}
type seasonCell struct {
	widget.BaseWidget
	poster  *canvas.Rectangle
	initial *canvas.Text
	title   *widget.Label
}

func StartApp() {
	a := app.New()
	w := a.NewWindow("Simple layout")
	w.Resize(fyne.NewSize(1200, 700))

	topBar := canvas.NewRectangle(color.NRGBA{
		R: 30, G: 30, B: 30, A: 255,
	})
	topBar.SetMinSize(fyne.NewSize(0, 40))

	sideBar := canvas.NewRectangle(color.NRGBA{
		R: 45, G: 45, B: 45, A: 255,
	})

	mainContent := canvas.NewRectangle(color.NRGBA{
		R: 25, G: 25, B: 25, A: 255,
	})

	split := container.NewHSplit(sideBar, mainContent)

	split.SetOffset(0.20)

	content := container.NewBorder(
		topBar,
		nil,
		nil,
		nil,
		split,
	)

	w.SetContent(content)
	w.ShowAndRun()

	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
