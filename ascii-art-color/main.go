package main

import (
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring> \"something\"")
}

func main() {
	// 1️⃣ parsing
	opts, ok := parseArgs(os.Args[1:])
	if !ok {
		printUsage()
		os.Exit(2)
	}

	// 2️⃣ load banner
	banner, err := NewBanner(opts.Font)
	if err != nil {
		fmt.Printf("Banner '%s' not found.\n", opts.Font)
		fmt.Println("Available banners:")

		files, _ := os.ReadDir("fonts")
		for _, f := range files {
			name := f.Name()
			if strings.HasSuffix(name, ".txt") {
				fmt.Println(" -", strings.TrimSuffix(name, ".txt"))
			}
		}
		return
	}

	// 3️⃣ render
	output, err := Render(banner, opts)

	if err != nil {
		fmt.Println("Render error:", err)
		os.Exit(1)
	}
	if opts.OutputFile != "" {
		err := os.WriteFile(opts.OutputFile, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			os.Exit(1)
		}
		fmt.Println("Output written to", opts.OutputFile)
		return

	}

	// 4️⃣ output
	fmt.Print(output)
}
