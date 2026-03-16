#!/bin/bash

echo "========================================"
echo " ASCII-ART COLOR — AUDIT TEST SUITE "
echo "========================================"
echo

run () {
  echo "▶ COMMAND:"
  echo "go run . $*"
  echo
  go run . "$@"
  echo
  echo "----------------------------------------"
  echo
}

# -------- Functional --------

# invalid usage
run --color red "banana"

# basic colors
run --color=red "hello world"
run --color=green "1 + 1 = 2"
run --color=yellow "(%&) ??"

# substring tests
run --color=red ello "hello"          # second until last
run --color=blue e "hello"             # single letter
run --color=cyan el "hello"            # two letters

# case-sensitive substring
run --color=orange GuYs "HeY GuYs"

# special characters
run --color=yellow "%" "(%&) ??"
run --color=purple "#" "###"

# uppercase / lowercase mix
run --color=cyan "Ab" "AbC xYz 123"

# -------- Bonus (color notations) --------

# HEX (must be quoted)
run --color="#ff0000" he "hello"
run --color="#00ffaa" lo "hello"

# RGB (quoted)
run --color="rgb(255,0,0)" h "hello"
run --color="rgb(0,255,170)" el "hello"

echo "======== END OF TESTS ========"
