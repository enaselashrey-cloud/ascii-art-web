package ascii

import (
	"strings"
)



func Render(banner *Banner, text string) (string, error){

	
	

	lines := strings.Split(text, `\n`)

	var result strings.Builder

	
//هيعمل سطر فاضي ف حالة وجود نيولاين
	for _, line := range lines {
		if line == "" {
			result.WriteByte('\n')
			continue
		}

		output := make([]string, 8)

		for _, ch := range line {
			rows, err := banner.GetChar(ch)
			if err != nil {
				return "", err
			}

			for i := range 8 {

				output[i] += rows[i]
			}
		}

		for _, asciiline := range output {
			result.WriteString(asciiline)
			result.WriteByte('\n')
		}
	}

	return result.String(), nil
}
