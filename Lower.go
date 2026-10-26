package main

import (
	"strings"
)

func MartiMinimus(sentence []string) []string {
	var result []string

	for _, s := range sentence {
		end := len(result) - 1
		if s == "(low)" {
			if end >= 0 {
				result[end] = strings.ToLower(result[end])
			}
		} else {
			result = append(result, s)
		}
	}
	return result
}
