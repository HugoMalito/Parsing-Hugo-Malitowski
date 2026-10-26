package main

import (
	"strings"
)

func LeMaxeur(sentence []string) []string {
	var result []string

	for _, s := range sentence {
		end := len(result) - 1
		if s == "(up)" {
			if end >= 0 {
				result[end] = strings.ToUpper(result[end])
			}
		} else {
			result = append(result, s)
		}
	}
	return result
}