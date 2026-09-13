package safemap

import (
	"strconv"
	"sync"
	"testing"
)

func TestSetGetLen(t *testing.T) {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			safeMap.Set("i"+strconv.Itoa(i), i)
			wg.Done()
		}()
	}

	wg.Wait()

	if safeMap.Len() != 100 {
		t.Error("Expected 100 elements")
	}

	for i := 0; i < 100; i++ {
		key := "i" + strconv.Itoa(i)
		value, exists := safeMap.Get(key)
		if !exists {
			t.Errorf("Element %v not fount", key)
		}

		if value != i {
			t.Errorf("Expected %d, but found %d", i, value)
		}
	}
}

func TestDeleteGetLen(t *testing.T) {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		safeMap.Set("i"+strconv.Itoa(i), i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				safeMap.Delete("i" + strconv.Itoa(i))
			}
			wg.Done()
		}()
	}

	wg.Wait()

	if safeMap.Len() != 90 {
		t.Error("Expected 90 elements")
	}

	for i := 10; i < 100; i++ {
		key := "i" + strconv.Itoa(i)
		value, exists := safeMap.Get(key)
		if !exists {
			t.Errorf("Element %v not fount", key)
		}

		if value != i {
			t.Errorf("Expected %d, but found %d", i, value)
		}
	}
}
