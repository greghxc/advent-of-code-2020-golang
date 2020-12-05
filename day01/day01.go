package main

import (
	"fmt"
	"strconv"

	"github.com/greghxc/aoc2020go/fileutil"
	"github.com/greghxc/aoc2020go/perfutil"
)

func partOne(input []int) (answer int) {
	return sub(2020, input)
}

func partTwo(input []int) (answer int) {
	for _, i := range input {
		s := sub(2020-i, input)
		if s > 0 {
			return i * s
		}
	}
	return
}

func sub(total int, vals []int) (product int) {
	seen := make(map[int]bool)
	for _, k := range vals {
		if seen[total-k] {
			return k * (total - k)
		}
		seen[k] = true
	}
	return
}

func makeIntList(strings []string) (ints []int) {
	for _, i := range strings {
		j, _ := strconv.Atoi(i)
		ints = append(ints, j)
	}
	return
}

func main() {
	lines := fileutil.GetFileAsStrings("./day01/day01.txt")

	times := 1
	fmt.Println("=== Day 01 ===")
	perfutil.Measure(times, func() string { return "Part One: " + strconv.Itoa(partOne(makeIntList(lines))) })
	perfutil.Measure(times, func() string { return "Part Two: " + strconv.Itoa(partTwo(makeIntList(lines))) })
}
