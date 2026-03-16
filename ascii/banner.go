package ascii

import (
	"errors"
	"os"
	"strings"
)

type Banner struct {
	font         map[rune][]string
}

func NewBanner(path string) (*Banner, error) {
	if !strings.HasSuffix(path, ".txt") {
		path += ".txt"
	}
	path = "fonts/" + path
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	for i := range lines {
		lines[i] = strings.TrimSuffix(lines[i], "\r")
	}
	for len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}

	b := &Banner{}
	b.font = make(map[rune][]string)

	//startASCII := 32
	//endASCII := 126

	for ascii := 32; ascii <= 126; ascii++ {
		start := (ascii - 32) * 9
		end := start + 8

		if end > len(lines) {
			return nil, errors.New("invalid banner file format: " + path)
		}
		b.font[rune(ascii)] = lines[start:end]
	}
	return b, nil

}

func (b *Banner) GetChar(r rune) ([]string, error) {
	if rows, ok := b.font[r]; ok {
		return rows, nil
	}
	return nil, errors.New("character not found in banner")

}
