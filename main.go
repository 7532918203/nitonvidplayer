package main

import (
	"fmt"
	"fyneVideoPlayer/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: animeplayer <library folder>")
		os.Exit(2)
	}
	library, error := parser.ScanLibrary(os.Args[1])
	if error != nil {
	}
	fmt.Println(library.Path)
}
