package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tile struct {
	number int
	data   []string
	edges  [4]string
}

func (t tile) New(s string) tile {
	tnew := tile{}
	s = strings.Split(s, ":")[0]
	n, _ := strconv.Atoi(strings.Split(s, " ")[1])
	tnew.number = n
	return tnew
}

func (t *tile) applyData(line string) {
	t.data = append(t.data, line)
}

func (t *tile) done() {
	t.edges[0] = t.data[0]
	t.edges[1] = getVerticalEdge(*t, 0)
	t.edges[2] = t.data[len(t.data)-1]
	t.edges[3] = getVerticalEdge(*t, len(t.edges[0])-1)
}

func getVerticalEdge(t tile, side int) string {
	var s string
	for _, v := range t.data {
		s += string(v[side])
	}
	return s
}

func (t tile) print() {
	fmt.Printf("tile: %d\n", t.number)
	for _, s := range t.data {
		fmt.Println(s)
	}
	for i, s := range t.edges {
		fmt.Println(i, s)
	}
}

func reverse(s string) string {
	var res string
	for i, _ := range s {
		res += string(s[len(s)-i-1])
	}
	return res
}
