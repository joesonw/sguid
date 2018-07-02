package sguid

import "testing"

func TestMemoryStorage(t *testing.T) {
	memory := NewMemoryStorage()

	count, err := memory.NewRange("0")
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Errorf("count %d is not %d", count, 0)
	}

	count, err = memory.NewRange("1")
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Errorf("count %d is not %d", count, 1)
	}

	count, err = memory.NewRange("0")
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Errorf("count %d is not %d", count, 2)
	}

	count, err = memory.NewRange("1")
	if err != nil {
		t.Error(err)
	}
	if count != 3 {
		t.Errorf("count %d is not %d", count, 3)
	}
}
