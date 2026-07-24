package validpalindrome

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1

	for left < right {
		for left < right && !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
			left++
		}

		for left < right && !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
			right--
		}

		if left < right && unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}

		left++
		right--
	}

	return true
}
