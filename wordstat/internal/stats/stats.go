package stats

import (
	"strings"
	"unicode"
)

func GetWords(input string) []string {
	runes := []rune(input)
	var words []string
	var word strings.Builder

	for _, r := range runes {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			if word.Len() > 0 {
				words = append(words, word.String())
			}
			word.Reset()
		} else {
			word.WriteRune(unicode.ToLower(r))
		}
	}

	if word.Len() > 0 {
		words = append(words, word.String())
	}

	return words
}

func GetWordToCount(words []string) map[string]int {
	wordToCount := make(map[string]int)

	for _, word := range words {
		wordToCount[word]++
	}

	return wordToCount
}