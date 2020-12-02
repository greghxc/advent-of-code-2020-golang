package main

import (
	"fmt"
	"strconv"

	"github.com/greghxc/aoc2020go/fileutil"
	"github.com/greghxc/aoc2020go/perfutil"
)

func partOne(input map[int]bool) (answer int) {
	return sub(2020, input)
}

func partTwo(input map[int]bool) (answer int) {
	for i := range input {
		s := sub(2020-i, input)
		if s > 0 {
			return i * s
		}
	}
	return
}

func sub(total int, seen map[int]bool) (product int) {
	for k := range seen {
		if seen[total-k] {
			return k * (total - k)
		}
	}
	return
}

func makeIntSet(strings []string) (ints map[int]bool) {
	ints = make(map[int]bool)
	for _, v := range strings {
		i, _ := strconv.Atoi(v)
		ints[i] = true
	}
	return
}

func main() {
	lines := fileutil.GetFileAsStrings("./day01/day01.txt")

	times := 1000
	fmt.Println("=== Day 01 ===")
	perfutil.Measure(times, func() string { return "Part One: " + strconv.Itoa(partOne(makeIntSet(lines))) })
	perfutil.Measure(times, func() string { return "Part Two: " + strconv.Itoa(partTwo(makeIntSet(lines))) })
}
