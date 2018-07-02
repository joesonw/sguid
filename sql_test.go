package sguid

import (
	"errors"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestSQLStorage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	sql := NewSQLStorage(db, "test")

	errMock := errors.New("mock error")

	mock.ExpectExec("INSERT INTO `test`").WithArgs("0").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO `test").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `test`").WithArgs("0").WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectExec("INSERT INTO `test`").WithArgs("1").WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectExec("INSERT INTO `test`").WithArgs("1").WillReturnError(errMock)

	count, err := sql.NewRange("0")
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Errorf("count %d is not %d", count, 0)
	}

	count, err = sql.NewRange("1")
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Errorf("count %d is not %d", count, 1)
	}

	count, err = sql.NewRange("0")
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Errorf("count %d is not %d", count, 2)
	}

	count, err = sql.NewRange("1")
	if err != nil {
		t.Error(err)
	}
	if count != 3 {
		t.Errorf("count %d is not %d", count, 3)
	}

	count, err = sql.NewRange("1")
	if err != errMock {
		t.Errorf("failed to get mock error")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
