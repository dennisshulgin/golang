package users

import (
	"encoding/json"
	"io"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Parse(input io.Reader) ([]User, error) {
	var users []User
	err := json.NewDecoder(input).Decode(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func CountActive(users []User) int {
	count := 0

	for _, user := range users {
		if user.Active {
			count++
		}
	}

	return count
}
func CountAdults(users []User) int {
	count := 0

	for _, user := range users {
		if user.Age >= 18 {
			count++
		}
	}

	return count
}
func AverageAge(users []User) float64 {
	result := 0.0
	count := len(users)

	if count == 0 {
		return result
	}

	for _, user := range users {
		result += float64(user.Age)
	}

	return result / float64(count)
}
