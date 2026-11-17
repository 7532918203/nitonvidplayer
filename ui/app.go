package ui

import (
	"fmt"
	"fyneVideoPlayer/model"
	"image/color"
	"regexp"
	"strings"

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

//helpers

func newShowCell(name string) *showCell {
	runes := []rune(name)
	initial := ""

	if len(runes) > 0 {
		initial = strings.ToUpper(string(runes[0]))
	}

	cell := &showCell{
		poster: canvas.NewRectangle(color.NRGBA{
			R: 55, G: 55, B: 55, A: 255,
		}),
		initial: canvas.NewText(initial, color.White),
		title:   widget.NewLabel(name),
	}

	cell.initial.TextSize = 28
	cell.initial.Alignment = fyne.TextAlignCenter
	cell.initial.TextStyle = fyne.TextStyle{Bold: true}

	cell.title.Alignment = fyne.TextAlignCenter
	cell.title.TextStyle = fyne.TextStyle{Bold: true}
	cell.title.Wrapping = fyne.TextWrapOff
	cell.title.Truncation = fyne.TextTruncateEllipsis

	cell.ExtendBaseWidget(cell)

	return cell
}

// CreateRenderer Override the CreateRenderer function to specify our own widget
func (c *showCell) CreateRenderer() fyne.WidgetRenderer {
	overlay := container.NewBorder(
		nil,
		c.title,
		nil,
		nil,
		c.initial,
	)

	content := container.NewStack(
		c.poster,
		overlay,
	)

	return widget.NewSimpleRenderer(content)
}

func cleanShowName(name string) string {
	//regex to remove all text within brackets []
	re := regexp.MustCompile(`\[[^]]*]`)
	cleaned := re.ReplaceAllString(name, "")

	return strings.TrimSpace(cleaned)
}

// file specific functions

func StartApp(library *model.Library) {
	a := app.New()
	w := a.NewWindow("Simple layout")
	w.Resize(fyne.NewSize(1200, 700))

	// TOP BAR
	topBar := canvas.NewRectangle(color.NRGBA{
		R: 30, G: 30, B: 30, A: 255,
	})
	topBar.SetMinSize(fyne.NewSize(0, 40))

	var showCells []fyne.CanvasObject

	for _, show := range library.Shows {
		showName := cleanShowName(show.Name)
		cell := newShowCell(showName)
		showCells = append(showCells, cell)
	}

	showGrid := container.NewGridWrap(
		fyne.NewSize(95, 100),
		showCells...,
	)

	showGridScroll := container.NewScroll(showGrid)

	sideBarHeader := widget.NewLabel("Library 1")
	sideBarHeader.Alignment = fyne.TextAlignCenter

	sideBar := container.NewBorder(
		sideBarHeader,
		nil,
		nil,
		nil,
		showGridScroll,
	)

	mainContent := canvas.NewRectangle(color.NRGBA{
		R: 25, G: 25, B: 25, A: 255,
	})

	split := container.NewHSplit(sideBar, mainContent)
	split.SetOffset(0.25)

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
