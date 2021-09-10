package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bfv/aoc2020-go/aocinput"
)

type Passport struct {
	fields map[string]string
}

var mustHave = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func (p *Passport) AddFields(kvpairs []string) {
	if p.fields == nil {
		p.fields = make(map[string]string)
	}
	for _, kvpair := range kvpairs {
		pair := strings.Split(kvpair, ":")
		p.fields[pair[0]] = pair[1]
	}
}

func (p Passport) IsValid() bool {
	for _, k := range mustHave {
		if _, ok := p.fields[k]; !ok {
			return false
		}
	}
	return true
}

func (p Passport) IsValidStrict() bool {

	ok := true

	for k, _ := range p.fields {

		v := p.fields[k]

		switch k {
		case "byr":
			year, _ := strconv.Atoi(v)
			ok = ok && (1920 <= year && year <= 2002)
		case "iyr":
			year, _ := strconv.Atoi(v)
			ok = ok && (2010 <= year && year <= 2020)
		case "eyr":
			year, _ := strconv.Atoi(v)
			ok = ok && (2020 <= year && year <= 2030)
		case "hgt":
			unit := regexp.MustCompile("[^a-zA-Z]+").ReplaceAllString(v, "")
			height, _ := strconv.Atoi(regexp.MustCompile("[^0-9]+").ReplaceAllString(v, ""))
			ok = ok && ((unit == "cm" && 150 <= height && height <= 193) || (unit == "in" && 59 <= height && height <= 76))
		case "hcl":
			match, _ := regexp.MatchString("^#([0-9a-f]{6})", v)
			ok = ok && match
		case "ecl":
			colors := map[string]struct{}{"amb": {}, "blu": {}, "brn": {}, "gry": {}, "grn": {}, "hzl": {}, "oth": {}}
			_, found := colors[v]
			ok = ok && found
		case "pid":
			_, err := strconv.Atoi(v)
			ok = ok && len(v) == 9 && err == nil
		}

		if !ok {
			break
		}
	}

	return ok
}

func main() {
	input := aocinput.GetStringSlice("input.txt")
	a, b := solve(input)

	fmt.Println("day4a:", a)
	fmt.Println("day4b:", b)
}

func solve(data []string) (int, int) {
	var validPassport, strictPassport int

	passports := getPassports(data)
	for _, passport := range passports {
		if passport.IsValid() {
			validPassport++
			if passport.IsValidStrict() {
				strictPassport++
			}
		}
	}

	return validPassport, strictPassport
}

func getPassports(data []string) []Passport {

	passports := make([]Passport, 0)

	passport := Passport{}
	for _, line := range data {
		if line == "" {
			passports = append(passports, passport)
			passport = Passport{}
		} else {
			passport.AddFields(strings.Split(line, " "))
		}
	}
	passports = append(passports, passport)

	return passports
}
