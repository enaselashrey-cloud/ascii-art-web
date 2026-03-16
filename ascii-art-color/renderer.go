package main

import (
	"strings"
)



func Render(banner *Banner, opts Options) (string, error){

	if opts.Text == "" {
		return "", nil
	}

	if opts.Text == `\n` {
		return "\n", nil
	}
	

	lines := strings.Split(opts.Text, `\n`)

	var result strings.Builder

	// تجهيز أدوات التلوين مرة واحدة
	colorCode := ""
	coloredPos := map[int]bool{}
	charIndex := 0

	if opts.Color != "" {
		if c, ok := getColorCode(opts.Color); ok {
			colorCode = c
		}
	}

	if opts.Color != "" && opts.Substring != "" {
		coloredPos = getColoredPosition(opts.Text, opts.Substring)
	}

	for _, line := range lines {
		if line == "" {
			result.WriteByte('\n')
			charIndex++ // علشان \n
			continue
		}

		output := make([]string, 8)

		for _, ch := range line {
			rows, err := banner.GetChar(ch)
			if err != nil {
				return "", err
			}

			for i := range 8 {

				//التلوين
				if colorCode != "" &&
					(opts.Substring == "" || coloredPos[charIndex]) {

					output[i] += colorCode + rows[i] + ansiReset
				} else {
					output[i] += rows[i]
				}
			}
			charIndex++
		}

		for _, asciiline := range output {
			result.WriteString(asciiline)
			result.WriteByte('\n')
		}
	}

	return result.String(), nil
}
