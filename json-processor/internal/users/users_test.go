package users

import (
	"strings"
	"testing"
)

func TestParseValid(t *testing.T) {
	input := strings.NewReader("[{\"id\": 1, \"name\": \"Denis\", \"age\": 28, \"active\": true}]")
	users, err := Parse(input)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	user := users[0]

	if user.ID != 1 || user.Active != true || user.Name != "Denis" || user.Age != 28 {
		t.Error("Invalid parse result")
	}
}

func TestParseInvalid(t *testing.T) {
	input := strings.NewReader("[{{\"id\": 1, \"name\": \"Denis\", \"age\": 28, \"active\": true}]")
	_, err := Parse(input)

	if err == nil {
		t.Error("Expected error")
	}
}

func TestCountActive(t *testing.T) {
	users := []User{
		{ID: 1, Name: "Denis", Age: 28, Active: true},
		{ID: 2, Name: "Anna", Age: 28, Active: false},
	}

	result := CountActive(users)

	if result != 1 {
		t.Errorf("Expected 1, but got %v", result)
	}
}

func TestCountAdults(t *testing.T) {
	users := []User{
		{ID: 1, Name: "Denis", Age: 28, Active: true},
		{ID: 2, Name: "Anna", Age: 28, Active: false},
	}

	result := CountAdults(users)

	if result != 2 {
		t.Errorf("Expected 2, but got %v", result)
	}
}

func TestAverageAge(t *testing.T) {
	users := []User{
		{ID: 1, Name: "Denis", Age: 28, Active: true},
		{ID: 2, Name: "Anna", Age: 28, Active: false},
	}

	result := AverageAge(users)

	if result != 28 {
		t.Errorf("Expected 28, but got %v", result)
	}
}

func TestAverageAgeEmptyArray(t *testing.T) {
	users := []User{}

	result := AverageAge(users)

	if result != 0 {
		t.Errorf("Expected 0, but got %v", result)
	}
}
