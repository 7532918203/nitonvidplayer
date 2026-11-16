package parser

import (
	"fmt"
	"fyneVideoPlayer/model"
)

func SiHei() {
	fmt.Println("hei på deg")
}

func ScanLibrary(path string) (*model.Library, error) {
	return &model.Library{Path: path}, nil
}
