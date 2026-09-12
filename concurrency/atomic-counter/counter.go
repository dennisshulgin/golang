package atomiccounter

import (
	"sync"
	"sync/atomic"
)

func UnsafeCounter(
	goroutines int,
	increments int,
) int64 {
	var counter int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Go(func() {
			for j := 0; j < increments; j++ {
				counter++
			}
		})
	}

	wg.Wait()

	return counter
}

func AtomicCounter(
	goroutines int,
	increments int,
) int64 {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Go(func() {
			for j := 0; j < increments; j++ {
				counter.Add(1)
			}
		})
	}

	wg.Wait()

	return counter.Load()
}
