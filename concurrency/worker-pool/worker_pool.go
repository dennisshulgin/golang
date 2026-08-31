package workerpool

import (
	"context"
	"sync"
)

func ProcessJob(
	ctx context.Context, 
	input <-chan int, 
	output chan int, 
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-input:
			if !ok {
				return
			}
			select {
				case <-ctx.Done():
					return 
				case output <- num * num:
			}
		}
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
	
	ctx, cancel := context.WithCancel(context.Background())

	input := make(chan int, workerCount)
	output := make(chan int, workerCount)
	
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go ProcessJob(ctx, input, output, &wg)
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
		cancel()
	} ()
	
	result := []int{}

	for out := range output {
		result = append(result, out)
	}

	return result
}