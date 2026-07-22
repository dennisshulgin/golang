package main

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

var Properties = []string{"host", "debug", "port"}
var BooleanValues = []string{"false", "true"}

var ErrInvalidPropertyLine = errors.New("Invalid property line")
var ErrInvalidHost = errors.New("Invalid host")
var ErrInvalidPort = errors.New("Invalid port")
var ErrInvalidDebug = errors.New("Invalid debug")
var ErrUnknownProperty = errors.New("Unknown property")

type Config struct {
	Host  string
	Port  int
	Debug bool
}

func ParseConfig(input string) (Config, error) {
	config := Config{}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		property := strings.Split(line, "=")
		if len(property) != 2 {
			return config, ErrInvalidPropertyLine
		}

		propertyError := validateProperty(property[0], property[1])

		if propertyError != nil {
			return config, propertyError
		}

		setProperty(&config, property[0], property[1])
	}

	return config, nil
}

func validateProperty(propertyName string, propertyValue string) error {
	if !slices.Contains(Properties, propertyName) {
		return ErrUnknownProperty
	}

	var propertyError error = nil

	if propertyName == "host" {
		propertyError = validateHost(propertyValue)
	}

	if propertyName == "debug" {
		propertyError = validateDebug(propertyValue)
	}

	if propertyName == "port" {
		propertyError = validatePort(propertyValue)
	}
	return propertyError
}

func validatePort(portAsString string) error {
	if len(portAsString) == 0 {
		return ErrInvalidPort
	}

	for i := 0; i < len(portAsString); i++ {
		if portAsString[i] < '0' || portAsString[i] > '9' {
			return ErrInvalidPort
		}
	}

	port, err := strconv.Atoi(portAsString)
	if err != nil || port < 1 || port > 65535 {
		return ErrInvalidPort
	}

	return nil
}

func validateDebug(debugAsString string) error {
	if !slices.Contains(BooleanValues, debugAsString) {
		return ErrInvalidDebug
	}

	return nil
}

func validateHost(hostAsString string) error {
	trimmedHost := strings.TrimSpace(hostAsString)
	if trimmedHost == "" {
		return ErrInvalidHost
	}

	return nil
}

func setProperty(config *Config, propertyName string, propertyValue string) {
	if propertyName == "host" {
		config.Host = propertyValue
	}
	if propertyName == "port" {
		port, _ := strconv.Atoi(propertyValue)
		config.Port = port
	}
	if propertyName == "debug" {
		debug, _ := strconv.ParseBool(propertyValue)
		config.Debug = debug
	}
}
