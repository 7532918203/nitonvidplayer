package parser

import (
	"fmt"
	"fyneVideoPlayer/model"
	"os"
	"path/filepath"
	"strings"
)

var videoExts = map[string]bool{
	".mp4": true, ".mkv": true, ".avi": true, ".mov": true,
	".webm": true, ".m4v": true, ".wmv": true, ".flv": true,
}

func isVideo(name string) bool {
	if strings.HasPrefix(name, ".") {
		return false
	}
	return videoExts[strings.ToLower(filepath.Ext(name))]
}

// ScanLibrary treats every directory directly under the specified path as a show.
func ScanLibrary(path string) (*model.Library, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	library := &model.Library{}

	for _, entry := range entries {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		show, err := scanShow(filepath.Join(path, entry.Name()), entry.Name())
		if err != nil {
			return nil, err
		}

		if len(show.Episodes) == 0 && len(show.Seasons) == 0 {
			continue
		}
		library.Shows = append(library.Shows, *show)
	}

	for _, show := range library.Shows {
		fmt.Println("Show:", show.Name)
		for _, ep := range show.Episodes {
			fmt.Println("  -", ep.Name)
		}
		for _, season := range show.Seasons {
			fmt.Println("  Season:", season.Name)
			for _, ep := range season.Episodes {
				fmt.Println("    -", ep.Name)
			}
		}
	}

	return library, nil
}

func scanShow(dir, name string) (*model.Show, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	show := &model.Show{Name: name}

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			if strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			season, err := scanSeason(fullPath, entry.Name())
			if err != nil {
				return nil, err
			}
			if len(season.Episodes) > 0 {
				show.Seasons = append(show.Seasons, season)
			}
		} else if isVideo(entry.Name()) {
			show.Episodes = append(show.Episodes, &model.Episode{
				Name: entry.Name(),
				Path: fullPath,
			})
		}
	}

	return show, nil
}

func scanSeason(dir, name string) (*model.Season, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	season := &model.Season{Name: name}

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			if strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			sub, err := scanSeason(fullPath, entry.Name())
			if err != nil {
				return nil, err
			}
			season.Episodes = append(season.Episodes, sub.Episodes...)
		} else if isVideo(entry.Name()) {
			season.Episodes = append(season.Episodes, &model.Episode{
				Name: entry.Name(),
				Path: fullPath,
			})
		}
	}

	return season, nil
}
