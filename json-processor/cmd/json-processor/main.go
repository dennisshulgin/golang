package main

import (
	"bufio"
	"fmt"
	"io"
	"json-processor/internal/users"
	"os"
)

func main() {
	args := os.Args

	var reader io.Reader

	if len(args) == 2 {
		file, err := os.Open(args[1])

		if err != nil {
			fmt.Printf("Open file error: %v", err)
			return
		}

		defer file.Close()

		reader = file
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	parsedUsers, err := users.Parse(reader)

	if err != nil {
		fmt.Printf("Parse input error: %v", err)
		return
	}

	fmt.Println("total users:", len(parsedUsers))
	fmt.Println("active users:", users.CountActive(parsedUsers))
	fmt.Println("adults:", users.CountAdults(parsedUsers))
	fmt.Println("average age:", users.AverageAge(parsedUsers))
}
