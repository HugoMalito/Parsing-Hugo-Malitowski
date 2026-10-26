package main

import (
	"strings"
)

func Capitalisme(sentence []string) []string {
	var result []string

	for _, c := range sentence {
		end := len(result) - 1
		if c == "(cap)" {
			if end >= 0 {
				letters := []rune(result[end])
				result[end] = strings.ToUpper(string(letters[0])) + strings.ToLower(string(letters[1:]))
			}
		} else {
			result = append(result, c)
		}
	}
	return result
}