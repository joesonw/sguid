package sguid

import "testing"

func TestCounter(t *testing.T) {
	config := NewConfig()
	counter := New(config)

	count, err := counter.Get()
	if err != nil {
		t.Error(err)
		return
	}

	if count != 0 {
		t.Errorf("count %d is not 0", count)
	}
}

func TestFragment(t *testing.T) {
	config := NewConfig().WithSize(2)
	counter := New(config)

	for i := 0; i < 5; i++ {
		count, err := counter.Get()
		if err != nil {
			t.Error(err)
			return
		}

		if count != int64(i) {
			t.Errorf("count %d is not %d", count, i)
		}
	}
}
