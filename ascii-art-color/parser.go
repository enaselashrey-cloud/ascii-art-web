package main

import "strings"

type Options struct {
	Text       string
	Font       string
	Color      string
	Substring  string
	OutputFile string
}

func parseArgs(args []string) (Options, bool) {
	opts := Options{
		Font: "standard",
	}

	if len(args) == 0 {
		return opts, false
	}

	var positional []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			opts.Color = strings.TrimPrefix(arg, "--color=")
			if opts.Color == "" {
				return opts, false
			}
		} else if strings.HasPrefix(arg, "--output=") {
			opts.OutputFile = strings.TrimPrefix(arg, "--output=")
			if opts.OutputFile == "" {
				return opts, false
			}

		} else if strings.HasPrefix(arg, "--") {
			return opts, false // flag غير معروف
		} else {
			positional = append(positional, arg)
		}
	}

	// لازم على الأقل text
	if len(positional) == 0 {
		return opts, false
	}

	if opts.Color != "" {
		if len(positional) == 1 {
			opts.Text = positional[0]
		}
		if len(positional) == 2 {
			opts.Substring = positional[0]
			opts.Text = positional[1]
		}
		if len(positional) == 3 {
			opts.Substring = positional[0]
			opts.Text = positional[1]
			opts.Font = positional[2]
		}
		return opts, true

	}

	opts.Text = positional[0]

	// آخر عنصر = font
	if opts.Color == "" && len(positional) >= 2 {
		opts.Font = positional[len(positional)-1]
	}

	

	return opts, true
}
