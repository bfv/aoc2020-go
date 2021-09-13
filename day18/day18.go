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
	sp    int
}

func (s *Stack) Push(char string) {
	s.chars = append(s.chars, char)
	(s.sp)++
}

func (s *Stack) Pop() string {
	ch := s.chars[s.sp-1]
	s.chars = s.chars[:s.sp-1]
	(s.sp)--
	return ch
}

func (s Stack) LastIsOperand() bool {
	if s.sp == 0 {
		return false
	}
	c := s.chars[s.sp-1]
	return c == "+" || c == "*"
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

	for _, expr := range data {
		answerA += calcExpressionA(expr)
		answerB += calcExpressionB(expr)
	}

	return answerA, answerB
}

func calcExpressionA(expr []string) int {

	stack := Stack{}

	for _, ch := range expr {
		if ch == "+" || ch == "*" || ch == "(" {
			stack.Push(ch)
		} else {
			if ch == ")" {
				ch = stack.Pop()
				stack.Pop() // remove the "("
			}
			processNumber(&stack, ch)
		}
	}

	res, _ := strconv.Atoi(stack.Pop())
	return res
}

func calcExpressionB(expr []string) int {
	return -1
}

func processNumber(stack *Stack, ch string) {

	if stack.LastIsOperand() {
		var res int
		thisInt, _ := strconv.Atoi(ch)
		operand := stack.Pop()
		otherInt, _ := strconv.Atoi(stack.Pop())
		if operand == "+" {
			res = thisInt + otherInt
		} else {
			res = thisInt * otherInt
		}
		ch = strconv.Itoa(res)
	}

	stack.Push(ch)
}
