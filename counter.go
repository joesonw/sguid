package sguid

import "sync"

// Counter sego counter
type Counter struct {
	storage Storage
	size    int64
	id      string
	idStart int64
	idStep  int64
	lock    *sync.Mutex
}

// New new server
func New(config *Config) *Counter {
	return &Counter{
		storage: config.storage,
		size:    config.size,
		id:      config.id,
		idStep:  0,
		idStart: -1,
		lock:    &sync.Mutex{},
	}
}

// Get get id
func (c *Counter) Get() (int64, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.idStep >= c.size {
		c.idStep = 0
		c.idStart = -1
	}

	if c.idStart == -1 {
		id, err := c.storage.NewRange(c.id)
		if err != nil {
			return 0, err
		}
		c.idStart = id * c.size
	}
	id := c.idStep + c.idStart
	c.idStep++
	return id, nil
}
