package goroutines

import (
	"fmt"
	"time"
	"sync"
)

func PrintNumber(worker int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("worker %d: %d \n", worker, i)
	}
}

func PrintNumbers() {
	goroutinesCount := 3
	numbersCount := 3

	for i := 1; i <= goroutinesCount; i++ {
		go PrintNumber(i, numbersCount)
	}

	time.Sleep(time.Second)
} 

func RunWorker(worker int, mu *sync.Mutex, wg *sync.WaitGroup, result *[]int) int {
	defer wg.Done()
	wg.Add(1)

	mu.Lock()
	*result = append(*result, worker)
	mu.Unlock()

	return worker
}

func RunWorkers(workerCount int) []int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result []int

	for i := 1; i <= workerCount; i++ {
		wg.Go(func() {
			RunWorker(i, &mu, &wg, &result)
		})
	}

	wg.Wait()

	return result
}


