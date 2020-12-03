package main

import (
	"fmt"
	"strconv"

	"github.com/greghxc/aoc2020go/fileutil"
	"github.com/greghxc/aoc2020go/perfutil"
)

type slope struct {
	right int
	down  int
}

func partOne(input []string) (trees int) {
	return sub(input, slope{3, 1})
}

func partTwo(input []string) (trees int) {
	slopes := []slope{slope{1, 1}, slope{3, 1}, slope{5, 1}, slope{7, 1}, slope{1, 2}}
	trees = 1
	for _, slope := range slopes {
		trees = trees * sub(input, slope)
	}
	return
}

func sub(input []string, slope slope) (trees int) {
	for i := 0; i*slope.down < len(input); i++ {
		row := i * slope.down
		column := (i * slope.right) % len(input[row])
		if input[row][column] == '#' {
			trees++
		}
	}
	return
}

func main() {
	lines := fileutil.GetFileAsStrings("./day03/day03.txt")

	times := 1
	fmt.Println("=== Day 03 ===")
	perfutil.Measure(times, func() string { return "Part One: " + strconv.Itoa(partOne(lines)) })
	perfutil.Measure(times, func() string { return "Part Two: " + strconv.Itoa(partTwo(lines)) })
}
