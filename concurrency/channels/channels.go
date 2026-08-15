package channels

import (
	"sync"
)

func GenerateSquares(numbers []int) []int {
	squares := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for i := 0; i < len(numbers); i++ {
		wg.Go(func() {
			calculateSquareAndSendToBuffer(numbers[i], squares)
		})
	}

	wg.Wait()
	close(squares)

	result := make([]int, 0, len(numbers))

	for square := range squares {
		result = append(result, square)
	}

	return result
}

func calculateSquareAndSendToBuffer(number int, squares chan int) {
	result := number * number
	squares <- result
}
