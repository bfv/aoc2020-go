package main

import (
	"fmt"

	"github.com/bfv/aoc2020-go/aocinput"
)

func main() {

	input := aocinput.GetInts("input.txt")
	a, b := solve(input)
	fmt.Println("day1a:", a)
	fmt.Println("day1b:", b)
}

func solve(input []int) (int, int) {
	answerA, answerB := -1, -1
loop:
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if answerA < 0 && (input[i]+input[j]) == 2020 {
				answerA = input[i] * input[j]
			} else if answerB < 0 && (input[i]+input[j]) < 2020 { // assumed: no zeros in input
				for k := 0; k < len(input); k++ {
					if (input[i] + input[j] + input[k]) == 2020 {
						answerB = input[i] * input[j] * input[k]
					}
				}
			}
			if answerA > -1 && answerB > -1 {
				break loop
			}
		}
	}
	return answerA, answerB
}
