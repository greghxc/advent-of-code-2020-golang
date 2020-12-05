package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/greghxc/aoc2020go/fileutil"
	"github.com/greghxc/aoc2020go/perfutil"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
}

func partOne(input []string) (valid int) {
	for _, line := range input {
		pp := passportBuilder(line)
		var results = []bool{
			pp.byr != "", pp.iyr != "", pp.eyr != "", pp.hgt != "", pp.hcl != "",
			pp.ecl != "", pp.pid != "",
		}
		if areAllTrue(results) {
			valid++
		}
	}
	return
}

func partTwo(input []string) (valid int) {
	for _, line := range input {
		pp := passportBuilder(line)
		var results = []bool{
			isIntWithinRange(pp.byr, 1920, 2002),
			isIntWithinRange(pp.iyr, 2010, 2020),
			isIntWithinRange(pp.eyr, 2020, 2030),
			isValidHeight(pp.hgt),
			isHexColor(pp.hcl),
			isEyeColor(pp.ecl),
			isDigits(pp.pid, 9),
		}
		if areAllTrue(results) {
			valid++
		}
	}
	return
}

func isIntWithinRange(input string, low int, high int) (result bool) {
	inputAsInt, _ := strconv.Atoi(input)
	return inputAsInt >= low && inputAsInt <= high
}

func isHexColor(input string) (result bool) {
	re := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	return re.MatchString(input)
}

func isEyeColor(input string) (result bool) {
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return re.MatchString(input)
}

func isDigits(input string, length int) (result bool) {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	return re.MatchString(input)
}

func isValidHeight(input string) (result bool) {
	re := regexp.MustCompile(`^(\d+)(cm|in)$`)
	if re.MatchString(input) {
		m := re.FindAllStringSubmatch(input, -1)
		switch measureUnit := m[0][2]; measureUnit {
		case "cm":
			return isIntWithinRange(m[0][1], 150, 193)
		case "in":
			return isIntWithinRange(m[0][1], 59, 76)
		default:
			return
		}
	} else {
		return
	}
}

func areAllTrue(input []bool) (result bool) {
	for _, b := range input {
		if !b {
			return
		}
	}
	return true
}

func passportBuilder(rawPassport string) (pp passport) {
	reFields := regexp.MustCompile("([a-z]+):([a-z0-9:#]+)")
	strippedLine := strings.ReplaceAll(rawPassport, "\n", " ")
	fieldMap := make(map[string]string)
	for _, field := range reFields.FindAllString(strippedLine, -1) {
		matches := reFields.FindAllStringSubmatch(field, -1)
		for _, match := range matches {
			fieldMap[match[1]] = match[2]
		}
	}
	pp = passport{
		byr: fieldMap["byr"],
		iyr: fieldMap["iyr"],
		eyr: fieldMap["eyr"],
		hgt: fieldMap["hgt"],
		hcl: fieldMap["hcl"],
		ecl: fieldMap["ecl"],
		pid: fieldMap["pid"],
	}
	return
}

func main() {
	lines := fileutil.GetFileAsStringsWithDelimiter("./day04/data/input.txt", "\n\n")

	times := 1
	fmt.Println("=== Day 04 ===")
	perfutil.Measure(times, func() string { return "Part One: " + strconv.Itoa(partOne(lines)) })
	perfutil.Measure(times, func() string { return "Part Two: " + strconv.Itoa(partTwo(lines)) })
}
