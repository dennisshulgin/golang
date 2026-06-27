package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

var Properties = []string{"host", "debug", "port"}
var booleanValues = []string{"false", "true"}

var ErrInvalidHost = errors.New("Invalid host")
var ErrInvalidPort = errors.New("Invalid port")
var ErrinvalidDebag = errors.New("Invalid debug")
var ErrUnknownProperty = errors.New("Unknown property")

func main() {
	file, err := os.Open("config.txt")
	hasErrors := false

	if err != nil {
		hasErrors = true
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "=")
		propertyName := splitted[0]
		propertyValue := splitted[1]
		err := validateProperty(propertyName, propertyValue)

		if err != nil {
			hasErrors = true
			fmt.Println(err.Error())
			break
		}
	}

	if scanner.Err(); err != nil {
		hasErrors = true
		fmt.Print(err)
	}

	if !hasErrors {
		fmt.Println("The file is valid!")
	}

}

func validateProperty(propertyName string, propertyValue string) error {
	if !slices.Contains(Properties, propertyName) {
		return ErrUnknownProperty
	}

	if propertyName == "host" {
		trimmedHost := strings.TrimSpace(propertyValue)
		if trimmedHost == "" {
			return ErrInvalidHost
		}
	}

	if propertyName == "debug" {
		if !slices.Contains(booleanValues, propertyValue) {
			return ErrinvalidDebag
		}
	}

	if propertyName == "port" {
		if len(propertyValue) == 0 {
			return ErrInvalidPort
		}

		for i := 0; i < len(propertyValue); i++ {
			if propertyValue[i] < '0' || propertyValue[i] > '9' {
				return ErrInvalidPort
			}
		}
	}
	return nil
}
