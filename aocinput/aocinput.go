package aocinput

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetStringArray(filename string) []string {
	return iterateFileLines(filename)
}

func GetInts(filename string) []int {

	lines := iterateFileLines((filename))
	ints := make([]int, 0, len(lines))

	for _, s := range lines {
		iv, _ := strconv.Atoi(s)
		ints = append(ints, iv)
	}

	return ints
}

func iterateFileLines(filename string) []string {

	lines := make([]string, 0, 16)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func GetScliceOfStringSlices(filename string) [][]string {

	lines := iterateFileLines((filename))
	stringSlice := make([][]string, 0, len(lines))

	for _, s := range lines {
		stringValues := strings.Split(s, " ")
		stringSlice = append(stringSlice, stringValues)
	}
	return stringSlice
}

func GetStringSlice(filename string) []string {

	lines := iterateFileLines((filename))

	stringSlice := make([]string, 0, len(lines))
	stringSlice = append(stringSlice, lines...)

	return stringSlice
}
