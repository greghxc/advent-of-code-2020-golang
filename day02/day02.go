package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/greghxc/aoc2020go/fileutil"
	"github.com/greghxc/aoc2020go/perfutil"
)

type passwordProfile struct {
	lower    int
	upper    int
	reqChar  string
	password string
}

func partOne(input []string) (validCount int) {
	for _, line := range input {
		pp := parsePasswordProfile(line)
		c := strings.Count(pp.password, pp.reqChar)
		if c >= pp.lower && c <= pp.upper {
			validCount++
		}
	}
	return
}

func partTwo(input []string) (validCount int) {
	for _, line := range input {
		pp := parsePasswordProfile(line)
		firstMatches := string(pp.password[pp.lower-1]) == pp.reqChar
		secondMatches := string(pp.password[pp.upper-1]) == pp.reqChar
		if firstMatches != secondMatches {
			validCount++
		}
	}
	return
}

func parsePasswordProfile(line string) (pp passwordProfile) {
	re := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)`)
	match := re.FindAllStringSubmatch(line, -1)
	lower, _ := strconv.Atoi(match[0][1])
	upper, _ := strconv.Atoi(match[0][2])
	return passwordProfile{
		lower:    lower,
		upper:    upper,
		reqChar:  match[0][3],
		password: match[0][4],
	}
}

func main() {
	lines := fileutil.GetFileAsStrings("./day02/day02.txt")

	times := 1
	fmt.Println("=== Day 02 ===")
	perfutil.Measure(times, func() string { return "Part One: " + strconv.Itoa(partOne(lines)) })
	perfutil.Measure(times, func() string { return "Part Two: " + strconv.Itoa(partTwo(lines)) })
}
