package main

import (
	"fmt"
	"os"
)

func main() {
	input, errReadFile := os.ReadFile("config.txt")

	if errReadFile != nil {
		fmt.Println(errReadFile)
		return
	}
	config, errConfig := ParseConfig(string(input))

	if errConfig != nil {
		fmt.Println("Config is invalid")
		fmt.Println(errConfig)
	} else {
		fmt.Println("Config is valid")
		fmt.Println(config)
	}
}
