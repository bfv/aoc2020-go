package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

type PwdRule struct {
	MinCount int
	MaxCount int
	Letter   string
}

func (p *PwdRule) ProcessInput(data []string) {
	counts := strings.Split(data[0], "-")
	p.MinCount, _ = strconv.Atoi(counts[0])
	p.MaxCount, _ = strconv.Atoi(counts[1])
	p.Letter = data[1][0:1]
}

func (p PwdRule) Validate(pwd string) bool {
	count := strings.Count(pwd, p.Letter)
	return p.MinCount <= count && count <= p.MaxCount
}

func main() {
	input := aocinput.GetScliceOfStringSlices("input.txt")
	a := solve(input)
	fmt.Println("day2a:", a)
}

func solve(input [][]string) int {
	var answer int

	for i := 0; i < len(input); i++ {
		rule, pwd := getRuleAndPwd(input[i])
		if rule.Validate(pwd) {
			answer++
		}
	}

	return answer
}

func getRuleAndPwd(data []string) (PwdRule, string) {

	rule := PwdRule{}
	rule.ProcessInput(data[:2])

	return rule, data[2]
}
