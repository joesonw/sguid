package sguid

import (
	"database/sql"
	"fmt"
)

// SQLStorage storage using sql databases
type SQLStorage struct {
	db    *sql.DB
	table string
}

// NewSQLStorage new SQLStorage
func NewSQLStorage(db *sql.DB, table string) *SQLStorage {
	return &SQLStorage{
		db,
		table,
	}
}

// NewRange Storage.NewRange
func (s *SQLStorage) NewRange(machineID string) (int64, error) {
	if s.db == nil {
		return 0, ErrStorageNotConnected
	}

	result, err := s.db.Exec(fmt.Sprintf("INSERT INTO `%s` (`machine_id`) VALUES (?)", s.table), machineID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
