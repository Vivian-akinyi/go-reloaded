package main

import (
	"strconv"
	"strings"
)

func ProcessUpper(text string) string {
	words := strings.Fields(text)
	for i := 1; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)

		}

	}
	return strings.Join(words, " ")

}

func ProcessUpperMultiple(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up," && i+1 < len(words) {
			countText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(countText)

			if err == nil {
				for j := 1; j <= count; j++ {
					if i-j >= 0 {
						words[i-j] = strings.ToUpper(words[i-j])
					}

				}

				words = append(words[:i], words[i+1:]...)
				words = append(words[:i], words[i+1:]...)
			}

		}

	}

	return strings.Join(words, " ")
}

func ProcessLower(text string) string {
	words := strings.Fields(text)
	for i := 1; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)

		}

	}
	return strings.Join(words, " ")
}

func ProcessCap(text string) string {
	words := strings.Fields(text)
	for i := 1; i < len(words); i++ {
		if words[i] == "(cap)" {
			word := words[i-1]
			words[i-1] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			words = append(words[:i], words[i+1:]...)

		}

	}
	return strings.Join(words, " ")
}
