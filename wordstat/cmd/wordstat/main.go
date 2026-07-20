package main

import (
	"example/wordstat/internal/stats"
	"fmt"
	"os"
)

func main() {
	reader, closeFunc, err := stats.GetInputReader(os.Args)
	defer closeFunc()

	input, err := stats.ReadAsString(reader)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	words := stats.GetWords(input)

	wordToCount := stats.GetWordToCount(words)

	fmt.Println("Words count: ", len(words))
	fmt.Println("Words: ", words)
	fmt.Println("Unique words count: ", len(wordToCount))
	fmt.Println("Unique words with counts: ", wordToCount)
}
