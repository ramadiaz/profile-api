package storages

import "sync"

type Memory struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewMemory() *Memory {
	return &Memory{
		data: make(map[string]interface{}),
	}
}

func (s *Memory) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Memory) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

func (s *Memory) Remove(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
