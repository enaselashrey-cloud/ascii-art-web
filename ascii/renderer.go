package ascii

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(font string) ([]string, error) {
	data, err := os.ReadFile("ascii/fonts/" + font + ".txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	return lines, nil
}

func RenderAscii(text string, lines []string) (string, error) {

	inputLines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var output strings.Builder

	for _, line := range inputLines {
		if line == "" {
			output.WriteByte('\n')
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range line {

				if ch < 32 || ch > 126 {
					return "", fmt.Errorf("unsupported character")
				}

				index := (int(ch)-32)*9 + row + 1

				if index >= len(lines) {
					return "", fmt.Errorf("invalid banner")
				}

				output.WriteString(lines[index])
			}
			output.WriteByte('\n')
		}
	}

	return output.String(), nil
}
