package main

import (
	"os/exec"
	"strings"
	"testing"
)

/* =========================
   parser.go – Unit Tests
========================= */

func TestParseArgs_TextOnly(t *testing.T) {
	args := []string{"Hello"}

	opts, ok := parseArgs(args)

	if !ok {
		t.Fatal("expected ok=true")
	}
	if opts.Text != "Hello" {
		t.Errorf("expected Text=Hello")
	}
	if opts.Font != "standard" {
		t.Errorf("expected default font")
	}
}

func TestParseArgs_TextAndFont(t *testing.T) {
	args := []string{"Hello", "standard"}

	opts, ok := parseArgs(args)

	if !ok {
		t.Fatal("expected ok=true")
	}
	if opts.Font != "standard" {
		t.Errorf("wrong font")
	}
}

func TestParseArgs_ColorOnly(t *testing.T) {
	args := []string{"--color=red", "Hello"}

	opts, ok := parseArgs(args)

	if !ok {
		t.Fatal("expected ok=true")
	}
	if opts.Color != "red" {
		t.Errorf("expected red")
	}
	if opts.Text != "Hello" {
		t.Errorf("wrong text")
	}
}

func TestParseArgs_ColorSubstring(t *testing.T) {
	args := []string{"--color=red", "lo", "Hello"}

	opts, ok := parseArgs(args)

	if !ok {
		t.Fatal("expected ok=true")
	}
	if opts.Substring != "lo" {
		t.Errorf("wrong substring")
	}
}

func TestParseArgs_Invalid(t *testing.T) {
	args := []string{"--color", "Hello"}

	_, ok := parseArgs(args)
	if ok {
		t.Fatal("expected ok=false")
	}
}

/* =========================
   banner.go – Unit Tests
========================= */

func TestNewBanner(t *testing.T) {
	b, err := NewBanner("standard")

	if err != nil {
		t.Fatalf("failed to load banner")
	}
	if b == nil {
		t.Fatal("banner is nil")
	}
}

func TestBannerGetChar(t *testing.T) {
	b, _ := NewBanner("standard")

	rows, err := b.GetChar('A')

	if err != nil {
		t.Fatalf("unexpected error")
	}
	if len(rows) != 8 {
		t.Errorf("expected 8 rows")
	}
}

func TestBannerInvalidChar(t *testing.T) {
	b, _ := NewBanner("standard")

	_, err := b.GetChar('❤')
	if err == nil {
		t.Fatal("expected error for invalid character")
	}
}

/* =========================
   color.go – Unit Tests
========================= */

func TestGetColorCode_Name(t *testing.T) {
	_, ok := getColorCode("red")
	if !ok {
		t.Fatal("expected valid named color")
	}
}

func TestGetColorCode_Hex(t *testing.T) {
	_, ok := getColorCode("#ff0000")
	if !ok {
		t.Fatal("expected valid hex color")
	}
}

func TestGetColorCode_RGB(t *testing.T) {
	_, ok := getColorCode("rgb(255,0,0)")
	if !ok {
		t.Fatal("expected valid rgb color")
	}
}

func TestGetColorCode_Invalid(t *testing.T) {
	_, ok := getColorCode("#ff")
	if ok {
		t.Fatal("expected invalid color")
	}
}

func TestGetColoredPosition(t *testing.T) {
	pos := getColoredPosition("Hello", "lo")

	if !pos[3] || !pos[4] {
		t.Fatal("substring positions incorrect")
	}
}

/* =========================
   renderer.go – Unit Tests
========================= */

func TestRendererSimple(t *testing.T) {
b, _ := NewBanner("standard")

opts := Options{Text: "A"}
out, err := Render(b, opts)


	if err != nil {
		t.Fatal(err)
	}
	if out == "" {
		t.Fatal("expected output")
	}
}

func TestRendererNewLine(t *testing.T) {
	b, _ := NewBanner("standard")

opts := Options{Text: "\\n"}
out, err := Render(b, opts)

	if err != nil {
		t.Fatal(err)
	}
	if out == "" {
		t.Fatal("expected multiline output")
	}
}

func TestRendererOnlyNewLine(t *testing.T) {
	b, _ := NewBanner("standard")

opts := Options{Text: "\\n"}
out, err := Render(b, opts)


	if err != nil {
		t.Fatal(err)
	}
	if out != "\n" {
		t.Fatalf("expected single newline")
	}
}

/* =========================
   Integration Tests
========================= */

func TestIntegration_BasicRun(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "Hello")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("program failed: %v\n%s", err, out)
	}

	if len(out) == 0 {
		t.Fatal("expected output, got empty")
	}
}

func TestIntegration_MultiLine(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "Hello\\nWorld")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("program failed: %v\n%s", err, out)
	}

	result := string(out)

	if !strings.Contains(result, "\n") {
		t.Fatal("expected multiline output")
	}
}

func TestIntegration_Color(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "--color=red", "Hello")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("program failed: %v\n%s", err, out)
	}

	if !strings.Contains(string(out), "\033[31m") {
		t.Fatal("expected ANSI red color code in output")
	}
}

func TestIntegration_InvalidBanner(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "Hello", "unknownfont")

	out, err := cmd.CombinedOutput()

	if err != nil {
		// main.go لا بيعمل os.Exit، ف ده طبيعي
		return
	}

	if !strings.Contains(string(out), "Available banners") {
		t.Fatal("expected banner error message")
	}
}
