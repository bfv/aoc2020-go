package main

import (
	"fmt"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

var tiles []tile = []tile{}

func main() {
	readTiles("input.txt")
	matchTiles()
	answerA := get20a()
	fmt.Printf("day20a: %v", answerA)
}

func matchTiles() {
	for i, t1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			t1.checkNeighbor(&tiles[j])
		}
	}
}

func get20a() int {
	answer := 1
	for _, t := range tiles {
		if len(t.neighbors) == 2 {
			answer *= t.number
		}
	}
	return answer
}

func readTiles(filename string) {
	var t tile

	input := aocinput.GetStringSlice(filename)
	for _, s := range input {
		if strings.HasPrefix(s, "Tile") {
			t = t.New(s)
			t.data = []string{}
		} else if s == "" {
			t.done()
			tiles = append(tiles, t)
		} else {
			t.applyData(s)
		}
	}
}
