package main

import (
	"strconv"
	"strings"
)

type tile struct {
	number int
	data   []string
}

func (t tile) New(s string) tile {
	tnew := tile{}
	s = strings.Split(s, ":")[0]
	n, _ := strconv.Atoi(strings.Split(s, " ")[1])
	tnew.number = n
	return tnew
}

func (t *tile) applyData(data []string) {
	t.data = data
}
