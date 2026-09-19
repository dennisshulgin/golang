package once

import (
	"sync"
	"testing"
)

func TestLoad(t *testing.T) {
	configLoader := NewConfigLoader()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Go(func() {
			configLoader.Load()
		})
	}
	wg.Wait()

	if configLoader.loadCount != 1 {
		t.Errorf("Expected 1 load count but got %d", configLoader.loadCount)
	}

	if configLoader.config.Host != "localhost" || configLoader.config.Port != 8080 {
		t.Error("Invalid config data")
	}
}
