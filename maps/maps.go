package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func ReadString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()	
}

func CountWords(s string) map[string]int {
	words := strings.Split(s, " ")
	wordToCount := make(map[string]int)	

	for _, word := range words {
		wordToCount[word]++
	}

	return wordToCount
}


func main() {
	text := ReadString()
	fmt.Println(CountWords(text))
}
