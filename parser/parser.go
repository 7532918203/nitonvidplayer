package parser

import (
	"fmt"
	"fyneVideoPlayer/model"
	"os"
)

func SiHei() {
	fmt.Println("hei på deg")
}

func ScanLibrary(path string) (*model.Library, error) {
	library := &model.Library{}
	entries, _ := os.ReadDir(path)
	for _, entry := range entries {
		library.Shows = append(library.Shows, entry.Name())
	}
	library.Path = path

	return library, nil
}
