package main

import (
	"testing"
)

func Test_partOne01(t *testing.T) {
	input := []string{
		"abc",
		"a\nb\nc",
		"ab\nac",
		"a\na\na\na",
		"b",
	}
	wantAnswer := 11
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partTwo01(t *testing.T) {
	input := []string{
		"abc",
		"a\nb\nc",
		"ab\nac",
		"a\na\na\na",
		"b",
	}
	wantAnswer := 6
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partTwo(input); gotAnswer != wantAnswer {
			t.Errorf("partTwo() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}
