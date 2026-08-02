package validparentheses

var openToClose map[rune]rune = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func IsValid(s string) bool {
	runes := []rune(s)

	stack := make([]rune, 0, len(s))

	for _, ch := range runes {
		if !isParenthese(ch) {
			return false
		}

		if isOpen(ch) {
			stack = append(stack, ch)
			continue
		}

		if isClose(ch) {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]

			if !isValidPair(last, ch) {
				return false
			}

			stack = stack[:len(stack)-1]

		}
	}

	return len(stack) == 0
}

func isOpen(ch rune) bool {
	if ch == '(' || ch == '{' || ch == '[' {
		return true
	}
	return false
}

func isClose(ch rune) bool {
	if ch == ')' || ch == '}' || ch == ']' {
		return true
	}
	return false
}

func isParenthese(ch rune) bool {
	if ch == '(' || ch == '{' || ch == '[' || ch == ')' || ch == '}' || ch == ']' {
		return true
	}
	return false
}

func isValidPair(open rune, close rune) bool {
	expectedClose, exists := openToClose[open]
	if exists && expectedClose == close {
		return true
	}
	return false
}
