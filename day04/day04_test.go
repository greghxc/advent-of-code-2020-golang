package main

import (
	"testing"

	"github.com/greghxc/aoc2020go/fileutil"
)

func Test_partOne01(t *testing.T) {
	input := fileutil.GetFileAsStringsWithDelimiter("../day04/data/test/p1.txt", "\n\n")
	wantAnswer := 2
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}
