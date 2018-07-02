package sguid

import (
	"sync"
)

// MemoryStorage in memory storage
type MemoryStorage struct {
	memory int64
	lock   *sync.Mutex
}

// NewMemoryStorage new memory storage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		memory: 0,
		lock:   &sync.Mutex{},
	}
}

// NewRange Storage.NewRange
func (s *MemoryStorage) NewRange(machineID string) (int64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	count := s.memory
	s.memory++
	return count, nil
}
