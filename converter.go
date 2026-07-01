package main

import (
	"strconv"
	"strings"
)

func ProcessHex(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			hexNumber := words[i-1]
			number, err := strconv.ParseInt(hexNumber, 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(number, 10)
				words = append(words[:i], words[i+1:]...)
			}
		}
	}

	return strings.Join(words, " ")
}

func ProcessBinary(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			binaryNumber := words[i-1]
			number, err := strconv.ParseInt(binaryNumber, 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(number, 10)
				words = append(words[:i], words[i+1:]...)

			}

		}

	}
	return strings.Join(words, " ")

}
