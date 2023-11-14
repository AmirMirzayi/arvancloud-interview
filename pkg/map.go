package pkg

import "sync"

type dataMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewMap() *dataMap {
	return &dataMap{
		data: make(map[string]int),
	}
}

func (m *dataMap) SetValue(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *dataMap) ExistedKey(key string) bool {
	if _, ok := m.data[key]; ok {
		return true
	}
	return false
}

func (m *dataMap) GetValue(key string) int {
	return m.data[key]
}

func (m *dataMap) GetData() map[string]int {
	return m.data
}
