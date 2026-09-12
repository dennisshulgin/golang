package atomiccounter

import (
	"testing"
)

func TestAtomicCounter(t *testing.T) {
	expected := 10000
	actual := AtomicCounter(10, 1000)

	if actual != int64(expected) {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func BenchmarkAtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AtomicCounter(10, 1000)
	}
}
