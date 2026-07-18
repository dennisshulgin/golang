package main

import (
	"bufio"
	"example/wordstat/internal/stats"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")

	input, err := stats.ReadAsString(reader)

	if err != io.EOF {
		return
	}

	words := stats.GetWords(input)

	wordToCount := stats.GetWordToCount(words)

	fmt.Println("Words count: ", len(words))
	fmt.Println("Words: ", words)
	fmt.Println("Unique words count: ", len(wordToCount))
	fmt.Println("Unique words with counts: ", wordToCount)
}
