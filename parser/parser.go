package parser

import (
	"fmt"
	"fyneVideoPlayer/model"
	"os"
	path2 "path"
)

func SiHei() {
	fmt.Println("hei på deg")
}

func ScanLibrary(path string) (*model.Library, error) {
	episodes, err := traverseDir(path)
	if err != nil {
		return nil, err
	}

	library := &model.Library{}

	fmt.Println("Episodes found:", len(episodes))

	for _, episode := range episodes {
		fmt.Println("- ", episode.Name)
	}
	return library, nil
}

func traverseDir(path string) ([]*model.Episode, error) {
	entries, _ := os.ReadDir(path)
	episodes := []*model.Episode{}
	episodeCount := 0

	for _, entry := range entries {
		fullPath := path2.Join(path, entry.Name())
		if entry.IsDir() {
			subEpisodes, _ := traverseDir(fullPath)

			episodes = append(episodes, subEpisodes...)
		} else {
			episodeCount++
			episodes = append(episodes, &model.Episode{entry.Name()})
		}
	}
	return episodes, nil
}
