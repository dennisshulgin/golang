package goroutines

import (
	"fmt"
	"sync"
)

func PrintNumber(worker int, count int, wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 1; i <= count; i++ {
		fmt.Printf("worker %d: %d \n", worker, i)
	}
	wg.Done()
}

func PrintNumbers() {
	goroutinesCount := 3
	numbersCount := 3
	var wg sync.WaitGroup

	for i := 1; i <= goroutinesCount; i++ {
		id := i
		go PrintNumber(id, numbersCount, &wg)
	}

	wg.Wait()
}

func RunWorker(worker int, mu *sync.Mutex, wg *sync.WaitGroup, result *[]int) {
	mu.Lock()
	*result = append(*result, worker)
	mu.Unlock()
}

func RunWorkers(workerCount int) []int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result []int

	for i := 1; i <= workerCount; i++ {
		id := i
		wg.Go(func() {
			RunWorker(id, &mu, &wg, &result)
		})
	}

	wg.Wait()
	return result
}
