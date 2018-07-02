package sguid

import "errors"

// storage errors
var (
	ErrStorageNotConnected = errors.New("storage is not connected")
)

// Storage storage interface
type Storage interface {
	NewRange(id string) (int64, error)
}
