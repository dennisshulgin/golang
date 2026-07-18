package main

import (
	"bufio"
	"example/wordstat/internal/stats"
	"fmt"
	"os"
	"io"
)

func main() {
	var reader io.Reader

	if len(os.Args) > 1 {
		path := os.Args[1]
		file, err := os.Open(path)

		if err != nil {
			return
		}
		defer file.Close()
		reader = file
	} else {
		fmt.Print("Input: ")
		reader = bufio.NewReader(os.Stdin)
	}

	input, err := stats.ReadAsString(reader)

	if err != nil {
		return
	}

	words := stats.GetWords(input)

	wordToCount := stats.GetWordToCount(words)

	fmt.Println("Words count: ", len(words))
	fmt.Println("Words: ", words)
	fmt.Println("Unique words count: ", len(wordToCount))
	fmt.Println("Unique words with counts: ", wordToCount)
}
