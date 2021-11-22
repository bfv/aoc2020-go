package main

import (
	"strconv"
	"strings"
)

// edges: top=0, left=1, bottom=2, right=3
type tile struct {
	number    int
	data      []string
	edges     [4]string
	neighbors map[int]bool
}

func (t tile) New(s string) tile {
	tnew := tile{}
	tnew.neighbors = make(map[int]bool)
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

func reverse(s string) string {
	var res string
	for i := range s {
		res += string(s[len(s)-i-1])
	}
	return res
}

func (t1 *tile) checkNeighbor(t2 *tile) {

	if t1.neighbors[t2.number] || t1.number == t2.number {
		return
	}

	match := false
	for i := 0; i < 4 && !match; i++ {
		for j := 0; j < 4 && !match; j++ {
			if t1.edges[i] == t2.edges[j] || t1.edges[i] == reverse(t2.edges[j]) {
				match = true
			}
		}
	}

	if match {
		t1.neighbors[t2.number] = true
		t2.neighbors[t1.number] = true
	}
}
