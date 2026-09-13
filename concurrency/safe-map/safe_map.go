package safemap

import "sync"

type SafeMap struct {
	mtx        sync.RWMutex
	keyToValue map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		mtx:        sync.RWMutex{},
		keyToValue: make(map[string]int),
	}
}

func (m *SafeMap) Set(key string, value int) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.keyToValue[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	value, exists := m.keyToValue[key]
	return value, exists
}

func (m *SafeMap) Delete(key string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.keyToValue, key)
}

func (m *SafeMap) Len() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return len(m.keyToValue)
}
