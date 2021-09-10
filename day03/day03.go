package main

import (
	"fmt"

	"github.com/bfv/aoc2020-go/aocinput"
)

func main() {
	input := aocinput.GetStringSlice("input.txt")
	a := solveA(input, 3, 1)
	b := solveB(input)
	fmt.Println("day3a:", a)
	fmt.Println("day3b:", b)
}

func solveA(input []string, dx int, dy int) int {
	return solve(input, dx, dy)
}

func solveB(input []string) int {
	return solve(input, 1, 1) * solve(input, 3, 1) * solve(input, 5, 1) * solve(input, 7, 1) * solve(input, 1, 2)
}

func solve(field []string, dx int, dy int) int {

	trees, x := 0, 0
	width := len(field[0])

	for y := 0; y < len(field); y += dy {
		if field[y][x:x+1] == "#" {
			trees++
		}
		x += dx
		x %= width
	}

	return trees
}
