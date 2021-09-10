package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

type PwdRule struct {
	i1     int
	i2     int
	letter string
}

func (p *PwdRule) ProcessInput(data []string) {
	counts := strings.Split(data[0], "-")
	p.i1, _ = strconv.Atoi(counts[0])
	p.i2, _ = strconv.Atoi(counts[1])
	p.letter = data[1][0:1]
}

func (p PwdRule) ValidateCount(pwd string) bool {
	count := strings.Count(pwd, p.letter)
	return p.i1 <= count && count <= p.i2
}

func (p PwdRule) ValidatePosition(pwd string) bool {
	return (pwd[p.i1-1:p.i1] == p.letter) != (pwd[p.i2-1:p.i2] == p.letter)
}

func main() {
	input := aocinput.GetScliceOfStringSlices("input.txt")
	a, b := solve(input)
	fmt.Println("day2a:", a)
	fmt.Println("day2b:", b)
}

func solve(input [][]string) (int, int) {
	var answerA, answerB int

	for i := 0; i < len(input); i++ {
		rule, pwd := getRuleAndPwd(input[i])
		if rule.ValidateCount(pwd) {
			answerA++
		}
		if rule.ValidatePosition(pwd) {
			answerB++
		}
	}

	return answerA, answerB
}

func getRuleAndPwd(data []string) (PwdRule, string) {

	rule := PwdRule{}
	rule.ProcessInput(data[:2])

	return rule, data[2]
}
