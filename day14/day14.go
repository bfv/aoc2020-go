package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

type BitMask struct {
	mask    string
	orMask  uint64
	andMask uint64
}

func (m *BitMask) Read(mask string) {
	m.mask = mask

	s := strings.ReplaceAll(mask, "X", "1")
	m.andMask, _ = strconv.ParseUint(s, 2, 64)

	s = strings.ReplaceAll(mask, "X", "0")
	m.orMask, _ = strconv.ParseUint(s, 2, 64)
}

func (m BitMask) Apply(number uint64) uint64 {
	number &= m.andMask
	number |= m.orMask
	return number
}

func (m BitMask) Convert(baseAddr int) BitMask {
	var ms string
	for i := len(m.mask) - 1; i >= 0; i-- {
		s := m.mask[i : i+1]
		var c string
		if s == "X" {
			c = "X"
		} else if s == "0" && (baseAddr&1 == 0) {
			c = "0"
		} else {
			c = "1"
		}
		ms = c + ms
		baseAddr >>= 1
	}
	newMask := BitMask{}
	newMask.Read(ms)

	return newMask
}

func (m BitMask) CalcAddrs(baseAddr int) []int {

	n := strings.Count(m.mask, "X")
	addrs := make([]int, 0)

	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		addr64, _ := strconv.ParseInt(applyToMask(m.mask, i), 2, 64)
		addr := int(addr64)
		addrs = append(addrs, addr)
	}
	return addrs
}

func main() {

	if true {
		input := aocinput.GetScliceOfStringSlices("input.txt", " ")
		a, b := solve(input)
		fmt.Println("day14a:", a)
		fmt.Println("day14b:", b)
	} else {
		m := BitMask{}
		m.mask = "X11X00X11"
		fmt.Println(m.CalcAddrs(0))
	}
}

func solve(data [][]string) (uint64, uint64) {

	var answerA, answerB uint64
	var mask BitMask

	mem := make(map[int]uint64)
	mem2 := make(map[int]uint64)

	for _, v := range data {

		if v[0] == "mask" {
			mask = BitMask{}
			mask.Read(v[2])
		} else {
			addr, val := processMemLine(v)
			mem[addr] = mask.Apply(val)

			mask2 := mask.Convert(addr)
			addrs := mask2.CalcAddrs(addr)
			for _, ma := range addrs {
				mem2[ma] = val
			}
		}
	}

	for _, v := range mem {
		answerA += v
	}

	for _, v := range mem2 {
		answerB += v
	}

	return answerA, answerB
}

func processMemLine(v []string) (addr int, val uint64) {
	val, _ = strconv.ParseUint(v[2], 10, 64)
	s := strings.ReplaceAll(v[0], "mem[", "")
	s = strings.ReplaceAll(s, "]", "")
	addr, _ = strconv.Atoi(s)
	return
}

func applyToMask(mask string, val int) string {

	pos := strings.LastIndex(mask, "X")
	for pos > -1 {
		mask = mask[:pos] + strconv.Itoa(val&1) + mask[pos+1:]
		val >>= 1
		pos = strings.LastIndex(mask, "X")
	}

	return mask
}
