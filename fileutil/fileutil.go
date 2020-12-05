package fileutil

import (
	"io/ioutil"
	"strings"
)

// GetFileAsStrings returns []string where each string is a line from file
func GetFileAsStrings(fileName string) []string {
	return GetFileAsStringsWithDelimiter(fileName, "\n")
}

// GetFileAsStringsWithDelimiter returns []string where each string is split by provided string
func GetFileAsStringsWithDelimiter(fileName string, delimiter string) []string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	return strings.Split(string(dat), delimiter)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
