package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/greghxc/aoc2020go/fileutil"
	"github.com/greghxc/aoc2020go/perfutil"
)

func partOne(input []string) (answer int) {
	for _, pass := range input {
		seat, _ := passToInt(pass)
		if seat > answer {
			answer = seat
		}
	}
	return
}

func partTwo(input []string) (answer int) {
	firstSeat, _ := passToInt(input[0])
	min := firstSeat
	max := firstSeat
	sum := firstSeat

	for _, pass := range input[1:] {
		seat, _ := passToInt(pass)
		if seat < min {
			min = seat
		}
		if seat > max {
			max = seat
		}
		sum = sum + seat
	}

	expectedTotalForRange := int((float32(max+min) / 2) * float32(max-min+1))
	return expectedTotalForRange - sum
}

func passToInt(pass string) (seat int, err error) {
	reOn := regexp.MustCompile(`[BR]`)
	reOff := regexp.MustCompile(`[FL]`)
	s := reOn.ReplaceAllString(pass, "1")
	s = reOff.ReplaceAllString(s, "0")
	i64, err := strconv.ParseInt(s, 2, 11)
	return int(i64), err
}

func main() {
	lines := fileutil.GetFileAsStringsWithDelimiter("./day05/data/input.txt", "\n")

	times := 1
	fmt.Println("=== Day 05 ===")
	perfutil.Measure(times, func() string { return "Part One: " + strconv.Itoa(partOne(lines)) })
	perfutil.Measure(times, func() string { return "Part Two: " + strconv.Itoa(partTwo(lines)) })
}
