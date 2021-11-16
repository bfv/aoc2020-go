package main

import (
	"fmt"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

var tiles []tile

func main() {
	var t tile
	var data []string

	input := aocinput.GetStringSlice("_input.txt")
	for _, s := range input {
		if strings.HasPrefix(s, "Tile") {
			t = t.New(s)
			data = []string{}
			fmt.Println("found tile", t.number)
		} else if s == "" {
			t.applyData(data)
			fmt.Println(data)
		} else {
			data = append(data, s)
		}
	}
}
