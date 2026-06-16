// package ascii

// import (
// 	"errors"
// 	"os"
// 	"strings"
// 	"sync"
// )

// var (
// 	ErrBannerNotFound       = errors.New("banner not found")
// 	ErrInvalidBanner        = errors.New("invalid banner")
// 	ErrUnsupportedCharacter = errors.New("unsupported character")
// )

// // bannerCache stores loaded banners to avoid repeated disk I/O
// var (
// 	bannerMutex sync.RWMutex
// 	bannerCache = make(map[string][]string)
// )

// func LoadBanner(font string) ([]string, error) {
// 	// Check cache first
// 	bannerMutex.RLock()
// 	if cached, exists := bannerCache[font]; exists {
// 		bannerMutex.RUnlock()
// 		return cached, nil
// 	}
// 	bannerMutex.RUnlock()

// 	// Load from disk
// 	data, err := os.ReadFile("ascii/fonts/" + font + ".txt")
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return nil, ErrBannerNotFound
// 		}
// 		return nil, err
// 	}

// 	lines := strings.Split(string(data), "\n")
// 	if len(lines) < 855 {
// 		return nil, ErrInvalidBanner
// 	}

// 	// Store in cache
// 	bannerMutex.Lock()
// 	bannerCache[font] = lines
// 	bannerMutex.Unlock()

// 	return lines, nil
// }

// func RenderAscii(text string, lines []string) (string, error) {

// 	inputLines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
// 	var output strings.Builder

// 	for _, line := range inputLines {
// 		if line == "" {
// 			output.WriteByte('\n')
// 			continue
// 		}

// 		for row := 0; row < 8; row++ {
// 			for _, ch := range line {

// 				if ch < 32 || ch > 126 {
// 					return "", ErrUnsupportedCharacter
// 				}

// 				index := (int(ch)-32)*9 + row + 1

// 				if index >= len(lines) {
// 					return "", ErrInvalidBanner
// 				}

// 				output.WriteString(lines[index])
// 			}
// 			output.WriteByte('\n')
// 		}
// 	}

// 	return output.String(), nil
// }

package ascii

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrBannerNotFound       = errors.New("banner not found")
	ErrInvalidBanner        = errors.New("invalid banner")
	ErrUnsupportedCharacter = errors.New("unsupported character")
)

func LoadBanner(font string) ([]string, error) {
	data, err := os.ReadFile("ascii/fonts/" + font + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrBannerNotFound
		}
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 855 {
		return nil, ErrInvalidBanner
	}

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
					return "", ErrUnsupportedCharacter
				}

				index := (int(ch)-32)*9 + row + 1
				fmt.Printf("%q\n", lines[(int('d')-32)*9+1])

				if index >= len(lines) {
					return "", ErrInvalidBanner
				}

				output.WriteString(lines[index])
			}
			output.WriteByte('\n')
		}
	}

	return output.String(), nil
}
