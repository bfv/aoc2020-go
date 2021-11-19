package main

import (
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

var tiles []tile

func main() {
	tiles = []tile{}
	readTiles()
}

func readTiles() {
	var t tile

	input := aocinput.GetStringSlice("_input.txt")
	for _, s := range input {
		if strings.HasPrefix(s, "Tile") {
			t = t.New(s)
			t.data = []string{}
		} else if s == "" {
			t.done()
			tiles = append(tiles, t)
			t.print()
		} else {
			t.applyData(s)
		}
	}
}
