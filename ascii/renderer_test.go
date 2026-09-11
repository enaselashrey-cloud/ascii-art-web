package ascii

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func chdir(t *testing.T, dir string) {
	t.Helper()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() returned an error: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("os.Chdir(%q) returned an error: %v", dir, err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})
}

func TestLoadBannerNotFound(t *testing.T) {
	chdir(t, "..")

	_, err := loadBanner("does-not-exist")
	if !errors.Is(err, ErrBannerNotFound) {
		t.Fatalf("loadBanner() error = %v; want ErrBannerNotFound", err)
	}
}

func TestLoadBannerInvalid(t *testing.T) {
	chdir(t, t.TempDir())

	if err := os.MkdirAll("ascii/fonts", 0o755); err != nil {
		t.Fatalf("os.MkdirAll() returned an error: %v", err)
	}
	if err := os.WriteFile("ascii/fonts/invalid.txt", []byte("invalid\n"), 0o644); err != nil {
		t.Fatalf("os.WriteFile() returned an error: %v", err)
	}

	_, err := loadBanner("invalid")
	if !errors.Is(err, ErrInvalidBanner) {
		t.Fatalf("loadBanner() error = %v; want ErrInvalidBanner", err)
	}
}

func TestRender(t *testing.T) {
	lines := make([]string, 855)
	for row := 0; row < 8; row++ {
		lines[(int('A')-32)*9+row+1] = "A" + string(rune('0'+row))
	}

	wantBlock := "A0\nA1\nA2\nA3\nA4\nA5\nA6\nA7\n"
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "single character", text: "A", want: wantBlock},
		{name: "empty line", text: "", want: "\n"},
		{name: "multiple lines", text: "A\nA", want: wantBlock + wantBlock},
		{name: "CRLF normalization", text: "A\r\nA", want: wantBlock + wantBlock},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := render(tt.text, lines)
			if err != nil {
				t.Fatalf("render() returned an error: %v", err)
			}
			if got != tt.want {
				t.Errorf("render() = %q; want %q", got, tt.want)
			}
		})
	}
}

func TestRenderUnsupportedCharacter(t *testing.T) {
	got, err := render("é", make([]string, 855))
	if !errors.Is(err, ErrUnsupportedCharacter) {
		t.Fatalf("render() error = %v; want ErrUnsupportedCharacter", err)
	}
	if strings.TrimSpace(got) != "" {
		t.Errorf("render() returned partial output %q on error", got)
	}
}

func TestOfficialAuditCases(t *testing.T) {
	chdir(t, "..")

	tests := []struct {
		name string
		font string
		text string
		want string
	}{
		{
			name: "standard braces and words",
			font: "standard",
			text: "{123}\n<Hello> (World)!",
			want: strings.Join([]string{
				"   __                     __",
				"  / /  _   ____    _____  \\ \\",
				" | |  / | |___ \\  |___ /   | |",
				"/ /   | |   __) |   |_ \\    \\ \\",
				"\\ \\   | |  / __/   ___) |   / /",
				" | |  |_| |_____| |____/   | |",
				"  \\_\\                     /_/",
				"",
				"   __  _    _          _   _          __            __ __          __                 _       _  __    _",
				"  / / | |  | |        | | | |         \\ \\          / / \\ \\        / /                | |     | | \\ \\  | |",
				" / /  | |__| |   ___  | | | |   ___    \\ \\        | |   \\ \\  /\\  / /    ___    _ __  | |   __| |  | | | |",
				"< <   |  __  |  / _ \\ | | | |  / _ \\    > >       | |    \\ \\/  \\/ /    / _ \\  | '__| | |  / _` |  | | | |",
				" \\ \\  | |  | | |  __/ | | | | | (_) |  / /        | |     \\  /\\  /    | (_) | | |    | | | (_| |  | | |_|",
				"  \\_\\ |_|  |_|  \\___| |_| |_|  \\___/  /_/         | |      \\/  \\/      \\___/  |_|    |_|  \\__,_|  | | (_)",
				"                                                   \\_\\                                           /_/",
				"",
			}, "\n") + "\n",
		},
		{
			name: "standard question marks",
			font: "standard",
			text: "123??",
			want: strings.Join([]string{
				"                     ___    ___",
				" _   ____    _____  |__ \\  |__ \\",
				"/ | |___ \\  |___ /     ) |    ) |",
				"| |   __) |   |_ \\    / /    / /",
				"| |  / __/   ___) |  |_|    |_|",
				"|_| |_____| |____/   (_)    (_)",
				"",
				"",
			}, "\n") + "\n",
		},
		{
			name: "shadow symbols",
			font: "shadow",
			text: "$% \"=",
			want: strings.Join([]string{
				"                        _|  _|",
				"  _|   _|_|    _|       _|  _|",
				"_|_|_| _|_|  _|                _|_|_|_|_|",
				"_|_|       _|",
				"  _|_|   _|  _|_|              _|_|_|_|_|",
				"_|_|_| _|    _|_|",
				"  _|",
				"",
			}, "\n") + "\n",
		},
		{
			name: "thinkertoy mixed characters",
			font: "thinkertoy",
			text: "123 T/fs#R",
			want: strings.Join([]string{
				"",
				"  0    --  o-o        o-O-o     o  o-o      | |  o--o",
				" /|   o  o    |         |      /   |       -O-O- |   |",
				"o |     /   oo          |     o   -O-  o-o  | |  O-Oo",
				"  |    /      |         |    /     |    \\  -O-O- |  \\",
				"o-o-o o--o o-o          o   o      o   o-o  | |  o   o",
				"",
				"",
			}, "\n") + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderAscii(tt.text, tt.font)
			if err != nil {
				t.Fatalf("RenderAscii(%q) returned an error: %v", tt.text, err)
			}

			got = trimTrailingSpaces(got)
			if got != tt.want {
				t.Errorf("RenderAscii(%q, %q) output differs from the official audit\ngot:\n%s\nwant:\n%s", tt.text, tt.font, got, tt.want)
			}
		})
	}
}

func trimTrailingSpaces(text string) string {
	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " ")
	}
	return strings.Join(lines, "\n")
}
