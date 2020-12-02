package fileutil

import (
	"io/ioutil"
	"strings"
)

// GetFileAsStrings returns []string where each string is a line from file
func GetFileAsStrings(fileName string) []string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	return strings.Split(string(dat), "\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
