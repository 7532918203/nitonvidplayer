package main

import (
	"fmt"
	"fyneVideoPlayer/parser"
	"fyneVideoPlayer/ui"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: animeplayer <library folder>")
		os.Exit(2)
	}
	library, _ := parser.ScanLibrary(os.Args[1])

	ui.StartApp(library)
}
