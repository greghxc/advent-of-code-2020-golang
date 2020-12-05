package main

import "testing"

func Test_partOne(t *testing.T) {
	input := makeIntList([]string{"1721", "979", "366", "299", "675", "1456"})
	wantAnswer := 514579
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partOneDuplicatsVals(t *testing.T) {
	input := makeIntList([]string{"1010", "1", "2019"})
	wantAnswer := 2019
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partOneDuplicatesValsOk(t *testing.T) {
	input := makeIntList([]string{"1010", "1010"})
	wantAnswer := 1020100
	t.Run("Part one provided example", func(t *testing.T) {
		if gotAnswer := partOne(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partTwo(t *testing.T) {
	input := makeIntList([]string{"1721", "979", "366", "299", "675", "1456"})
	wantAnswer := 241861950
	t.Run("Part two provided example", func(t *testing.T) {
		if gotAnswer := partTwo(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partTwoDuplicates(t *testing.T) {
	input := makeIntList([]string{"1000", "20", "40", "80", "1900"})
	wantAnswer := 6080000
	t.Run("Part two provided example", func(t *testing.T) {
		if gotAnswer := partTwo(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}

func Test_partTwoDuplicatesOk(t *testing.T) {
	input := makeIntList([]string{"1000", "1000", "20"})
	wantAnswer := 20000000
	t.Run("Part two provided example", func(t *testing.T) {
		if gotAnswer := partTwo(input); gotAnswer != wantAnswer {
			t.Errorf("partOne() = %v, want %v", gotAnswer, wantAnswer)
		}
	})
}
