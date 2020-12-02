package main

import "testing"

func Test_partOne(t *testing.T) {
	input := makeIntSet([]string{"1721", "979", "366", "299", "675", "1456"})
	wantAnswer := 514579
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partTwo(t *testing.T) {
	input := makeIntSet([]string{"1721", "979", "366", "299", "675", "1456"})
	wantAnswer := 241861950
	t.Run("Part two provided example", func(t *testing.T) {
		if gotAnswer := partTwo(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}
