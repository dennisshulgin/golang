package main

import (
	"bufio"
	"fmt"
	"os"
	"example/wordstat/internal/stats"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input: ")
	scanner.Scan()

	input := scanner.Text()

	words := stats.GetWords(input)
	wordToCount := stats.GetWordToCount(words)

	fmt.Println("Words count: ", len(words))
	fmt.Println("Words: ", words)
	fmt.Println("Unique words count: ", len(wordToCount))
	fmt.Println("Unique words with counts: ", wordToCount)

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error: ", err)
	}
}