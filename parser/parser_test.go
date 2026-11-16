package parser

import "testing"

func TestParseFilename(t *testing.T) {
	tests := []struct {
		input   string
		title   string
		episode int
	}{
		{"Frieren - 05", "Frieren", 5},
		{"Frieren S01E05", "Frieren", 5},
		{"Frieren Episode 05", "Frieren", 5},
		{"Frieren 1x05", "Frieren", 5},
	}

	for _, tt := range tests {
		result := parseFilename(tt.input)

		if result.Title != tt.title {
			t.Errorf(
				"%s: expected title %q, got %q",
				tt.input, tt.title, result.Title,
			)
		}

		if result.Episode != tt.episode {
			t.Errorf(
				"%s: expected episode %d, got %d",
				tt.input, tt.episode, result.Episode,
			)
		}
	}
}
