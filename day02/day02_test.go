package main

import "testing"

func Test_partOne(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	wantAnswer := 2
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partTwo(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	wantAnswer := 1
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partTwo(input); gotAnswer != wantAnswer {
			t.Errorf("partTwo() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}
