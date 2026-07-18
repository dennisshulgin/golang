package stats

import (
	"io"
	"strings"
	"unicode"
)

func ReadAsString(input io.Reader) (string, error) {
	buffer := make([]byte, 8)
	var inputString strings.Builder

	for {
		count, err := input.Read(buffer)

		if count > 0 {
			for i, ch := range buffer[:count] {
				if ch == '\n' {
					inputString.Write(buffer[:i])
					return inputString.String(), nil
				}
			}
		}

		if err != nil {
			if err == io.EOF {
				return inputString.String(), nil
			}

			return inputString.String(), err
		}
	}
}

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
