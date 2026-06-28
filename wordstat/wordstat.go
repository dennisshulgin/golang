package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input: ")
	scanner.Scan()

	input := scanner.Text()

	words := getWords(input)
	wordToCount := getWordToCount(words)

	fmt.Println("Words count: ", len(words))
	fmt.Println("Words: ", words)
	fmt.Println("Unique words count: ", len(wordToCount))
	fmt.Println("Unique words with counts: ", wordToCount)

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error: ", err)
	}
}

func getWords(input string) []string {
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

func getWordToCount(words []string) map[string]int {
	wordToCount := make(map[string]int)

	for _, word := range words {
		wordToCount[word]++
	}

	return wordToCount
}
