package main

import (
	"fmt"
	"strconv"

	"github.com/greghxc/aoc2020go/fileutil"
	"github.com/greghxc/aoc2020go/perfutil"
)

func partOne(input []string) (count int) {
	for _, group := range input {
		var charMap = make(map[rune]bool)
		for _, c := range group {
			if c != '\n' {
				charMap[c] = true
			}
		}
		fmt.Println(charMap)
		count += len(charMap)
	}
	return
}

func partTwo(input []string) (count int) {
	for _, group := range input {
		var charMap = make(map[rune]int)
		for _, c := range group {
			charMap[c] = charMap[c] + 1
		}
		members := charMap['\n'] + 1
		for k, v := range charMap {
			if k != '\n' && v == members {
				count++
			}
		}
	}
	return
}

func main() {
	lines := fileutil.GetFileAsStringsWithDelimiter("./day06/data/input.txt", "\n\n")

	times := 1
	fmt.Println("=== Day 05 ===")
	perfutil.Measure(times, func() string { return "Part One: " + strconv.Itoa(partOne(lines)) })
	perfutil.Measure(times, func() string { return "Part Two: " + strconv.Itoa(partTwo(lines)) })
}
