package pkg

import "sync"

type DataMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewDataMap() *DataMap {
	return &DataMap{
		data: make(map[string]int),
	}
}

func (m *DataMap) SetValue(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *DataMap) ExistedKey(key string) bool {
	if _, ok := m.data[key]; ok {
		return true
	}
	return false
}

func (m *DataMap) GetValue(key string) int {
	return m.data[key]
}

func (m *DataMap) GetData() map[string]int {
	return m.data
}
