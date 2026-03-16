# 🎨 ASCII ART COLOR — GO PROJECT

A clean and modular **ASCII Art CLI tool** written in **Go**.  
Converts text into ASCII art using banner files with full support for  
**File System fonts**, **multi-line input**, and **ANSI color output**.

---

## 📑 Table of Contents
- [Overview](#overview)
- [Usage](#usage)
- [Features](#features)
- [Color System](#color-system)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Roadmap](#roadmap)
- [Author](#author)

---

## Overview
This project generates ASCII Art from text using banner files stored on disk.
It is designed to be simple, readable, and easily extensible.

Key goals:
- Clean Go code
- Modular design
- Clear separation of responsibilities
- Easy feature expansion (color, justify, output, etc.)

---

## Usage

Basic:
go run . "Hello"
go run . "Hello" standard
go run . "Hello\nWorld"

With color:
go run . --color=red "Hello"
go run . --color=#ff0000 "Hello"
go run . --color=rgb(0,255,0) "Hello"

Color substring:
go run . --color=blue lo "Hello"

---

## Features

- ASCII rendering for characters 32 → 126
- 8-line ASCII output per character
- Processes text row-by-row
- Supports escaped new lines (\n)
- File System based fonts (fs)
- Automatic font discovery
- ANSI color output
- Substring coloring
- Safe ANSI reset
- Robust error handling

---

## Color System

Supported color formats:
- Named colors (red, green, blue, yellow, etc.)
- HEX format (#ff0000)
- RGB format (rgb(255,0,0))

Color is parsed once and applied during rendering.
Substring coloring is handled using indexed character positions for accuracy.

---

## Architecture

parser.go  
Handles CLI argument parsing and validation.

banner.go  
Loads banner files from the file system and maps runes to ASCII rows.

renderer.go  
Core rendering engine responsible for:
- New line handling
- ASCII composition
- Color application

color.go  
Parses ANSI color formats and calculates substring positions.

main.go  
Program entry point that controls execution flow.

---

## Project Structure

ascii-art-color/
├── main.go
├── parser.go
├── renderer.go
├── banner.go
├── color.go
├── main_test.go
├── fonts/
│   └── standard.txt
└── go.mod

---

## Testing

Run all tests:
go test ./...

Tests include:
- Argument parsing
- Banner loading
- Renderer behavior
- Color parsing
- Full program integration

---

## Roadmap

- ASCII Art Core ✅
- File System (fs) ✅
- Color Output ✅
- Substring Coloring ✅
- Justify / Alignment ⏳
- Output to File ⏳
- Reverse Mode ⏳

---

## Author

Enas Essam
