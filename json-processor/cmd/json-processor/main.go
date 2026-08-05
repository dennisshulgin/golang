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

	users, err := users.Parse(reader)

	if err != nil {
		fmt.Printf("Parse input error: %v", err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
