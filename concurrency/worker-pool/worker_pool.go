package workerpool

import (
	"sync"
)

func ProcessJob(
	input <-chan int,
	output chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for num := range input {
		output <- num * num
	}
}

func ProcessJobs(
	jobs []int,
	workerCount int,
) []int {
	if workerCount <= 0 {
		return []int{}
	}

	var wg sync.WaitGroup

	input := make(chan int, workerCount)
	output := make(chan int, workerCount)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go ProcessJob(input, output, &wg)
	}

	go func() {
		for _, job := range jobs {
			input <- job
		}
		close(input)
	}()

	go func() {
		wg.Wait()
		close(output)
	}()

	result := []int{}

	for out := range output {
		result = append(result, out)
	}

	return result
}
