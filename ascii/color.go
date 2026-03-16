package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

var ansiColors = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"white":  "\033[37m",
	"orange": "\033[38;5;214m",
}

const ansiReset = "\033[0m"

func getColorCode(input string) (string, bool) {

	// 1️⃣ اسم لون
	if code, ok := ansiColors[input]; ok {
		return code, true
	}

	// 2️⃣ لون بالأرقام (#ff0000)
	if strings.HasPrefix(input, "#") {

		if len(input) != 7 {
			return "", false
		}

		r, err1 := strconv.ParseInt(input[1:3], 16, 0)
		g, err2 := strconv.ParseInt(input[3:5], 16, 0)
		b, err3 := strconv.ParseInt(input[5:7], 16, 0)

		if err1 != nil || err2 != nil || err3 != nil {
			return "", false
		}

		return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), true
	}
	// 3️⃣ لون بصيغة rgb(…)
if strings.HasPrefix(input, "rgb(") && strings.HasSuffix(input, ")") {

	values := strings.Split(input[4:len(input)-1], ",")
	if len(values) != 3 {
		return "", false
	}

	r, err1 := strconv.Atoi(strings.TrimSpace(values[0]))
	g, err2 := strconv.Atoi(strings.TrimSpace(values[1]))
	b, err3 := strconv.Atoi(strings.TrimSpace(values[2]))

	if err1 != nil || err2 != nil || err3 != nil {
		return "", false
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), true
}


	return "", false
}

func getColoredPosition(text, substring string) map[int]bool {
	pos := make(map[int]bool)
	if substring == "" {
		return pos
	}
	for i := 0; i <= len(text)-len(substring); i++ {
		if text[i:i+len(substring)] == substring {
			for j := 0; j < len(substring); j++ {
				pos[i+j] = true
			}
		}
	}
	return pos
}
