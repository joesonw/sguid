package sguid

import (
	"github.com/satori/go.uuid"
)

// Config config
type Config struct {
	id      string
	storage Storage
	size    int64
}

// NewConfig new config
func NewConfig() *Config {
	id := uuid.NewV4()
	return &Config{
		storage: NewMemoryStorage(),
		size:    100,
		id: id.String(),
	}
}

// WithStorage set storage
func (c *Config) WithStorage(storage Storage) *Config {
	c.storage = storage
	return c
}

// WithSize set fragment size
func (c *Config) WithSize(size int64) *Config {
	c.size = size
	return c
}

// WithID set machien id
func (c *Config) WithID(id string) *Config {
	c.id = id
	return c
}
