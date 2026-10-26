package main

import "strings"

func Anal_yse(sentence []string) []string {
	result := make([]string, len(sentence))
	copy(result, sentence)

	for i := 0; i < len(result)-1; i++ {
		if result[i] == "a" || result[i] == "A" {
			next := []rune(strings.ToLower(result[i+1]))
			if len(next) > 0 && strings.ContainsRune("aeiouh", next[0]) {
				result[i] += "n" // a -> an, A -> An
			}
		}
	}
	return result
}