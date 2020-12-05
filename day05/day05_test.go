package main

import (
	"testing"
)

func Test_partOne01(t *testing.T) {
	input := []string{"FBFBBFFRLR"}
	wantAnswer := 357
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partOne02(t *testing.T) {
	input := []string{"BFFFBBFRRR"}
	wantAnswer := 567
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partOne03(t *testing.T) {
	input := []string{"FFFBBBFRRR"}
	wantAnswer := 119
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partOne04(t *testing.T) {
	input := []string{"BBFFBBFRLL"}
	wantAnswer := 820
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}
