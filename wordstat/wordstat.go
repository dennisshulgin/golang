package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var delimiters = map[byte]bool{
	' ': true,
	'?': true,
	'.': true,
	'!': true,
}

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
	var words []string
	var word strings.Builder
	length := len(input)

	for i := 0; i < length; i++ {
		if delimiters[input[i]] {
			if word.Len() > 0 {
				words = append(words, word.String())
			}
			word.Reset()
		} else {

			if input[i] >= 65 && input[i] <= 90 {
				word.WriteByte(input[i] + 32)
			} else {
				word.WriteByte(input[i])
			}
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
