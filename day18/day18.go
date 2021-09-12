package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

func main() {

	input := aocinput.GetStringSlice("_input.txt")
	data := processInput(input)
	a, b := solve(data)
	fmt.Println("day18a:", a)
	fmt.Println("day18b:", b)
}

type Stack struct {
	chars []string
	pos int
}

func (s *Stack) Push(char string) {
	s = append(s, char)
	pos++
}

func (s Stack) LastIsOperand() bool {
	c := s.chars[pos-1]
	return c == "+" || c == "*"
}

func (s Stack) LastIsNumber() bool {
	i, err := strconv.Atoi(s.chars[pos-1])
	return err == nil
}

func processInput(input []string) [][]string {
	var data [][]string

	for _, v := range input {
		v = strings.ReplaceAll(v, " ", "")
		s := strings.Split(v, "")
		data = append(data, s)
	}
	return data
}

func solve(data [][]string) (int, int) {
	var answerA, answerB int

	for a, expr := range data {
		parseExpression(expr[a+1:])
	}

	fmt.Println(data)
	return answerA, answerB
}

func parseExpression(expr []string) int {
	var res int
	var lastNumber int
	var lastOperand string

	for _, char := range expr {
		switch char {
		case "(":
			res = 
		}
	}
}