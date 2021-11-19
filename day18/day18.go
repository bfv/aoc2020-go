package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

func main() {

	input := aocinput.GetStringSlice("input.txt")
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
	var ch string

	if s.sp > 0 {
		ch = s.chars[s.sp-1]
		s.chars = s.chars[:s.sp-1]
		(s.sp)--
	}
	return ch
}

func (s Stack) LastIsOperand() bool {
	if s.sp == 0 {
		return false
	}
	c := s.chars[s.sp-1]
	return c == "+" || c == "*"
}

func (s Stack) IsEmpty() bool {
	return s.sp == 0
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

	for i, expr := range data {
		answerA += calcExpressionA(expr)
		fmt.Println("a", i)
		answerB += calcExpressionB(expr)
		fmt.Println("b", i)
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
			processNumberA(&stack, ch)
		}
	}

	res, _ := strconv.Atoi(stack.Pop())
	return res
}

func processNumberA(stack *Stack, ch string) {

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

func calcExpressionB(expr []string) int {
	stack := Stack{}

	for _, ch := range expr {
		force := false
		if ch == "+" || ch == "*" || ch == "(" {
			stack.Push(ch)
		} else {
			if ch == ")" {
				ch = stack.Pop()
				force = true
			}
			processNumberB(&stack, ch, force)
		}
	}

	// at this point there's just multiplication left
	for stack.sp > 1 {
		ch := stack.Pop()
		processNumberB(&stack, ch, true)
	}

	res, _ := strconv.Atoi(stack.Pop())
	return res
}

func processNumberB(stack *Stack, ch string, force bool) {

	if stack.LastIsOperand() {
		operand := stack.Pop()
		if operand == "+" || force {
			res := applyOperandB(stack.Pop(), operand, ch)
			if force && !(stack.IsEmpty()) {
				popped := stack.Pop()
				if popped != "(" {
					stack.Push(popped)
				}
			}
			if force {
				processNumberB(stack, res, false)
			} else {
				stack.Push(res)
			}
		} else {
			stack.Push(operand) // push * back on the stack
			stack.Push(ch)
		}
	} else {
		if !stack.IsEmpty() {
			popped := stack.Pop()
			if popped != "(" {
				stack.Push(popped)
			}
		}
		stack.Push(ch)
	}

}

func applyOperandB(a string, operand string, b string) string {
	var res int
	iA, _ := strconv.Atoi(a)
	iB, _ := strconv.Atoi(b)
	if operand == "+" {
		res = iA + iB
	} else {
		res = iA * iB
	}
	return strconv.Itoa(res)
}
