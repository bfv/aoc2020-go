package main

import "fmt"

type edge struct {
	data  string
	my    int
	other int // number
}

func (t tile) print() {
	fmt.Printf("tile: %d, neighbors: %v\n", t.number, t.neighbors)
}

func (t tile) printGrid() {
	for _, s := range t.data {
		fmt.Println(s)
	}
}

func (t tile) printEdges() {
	for i, s := range t.edges {
		fmt.Println(i, s)
	}
}

func (t *tile) flipX() {
	for i := 0; i < len(t.data)/2; i++ {
		ti := t.data[i]
		tn := t.data[len(t.data)-1-i]
		t.data[i], t.data[len(t.data)-1-i] = tn, ti
	}
	t.done()
}

func (t *tile) flipY() {
	for i, _ := range t.data {
		t.data[i] = reverse(t.data[i])
	}
	t.done()
}

func (t *tile) rotate() {
}
