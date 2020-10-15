package storage

import (
	"testing"
)

func TestMemory(t *testing.T) {
	td := []struct {
		Key   string
		Exist bool
	}{
		{Key: "first", Exist: true},
		{Key: "second", Exist: true},
		{Key: "some_one", Exist: false},
	}

	m := NewMemory()
	for _, d := range td {
		if d.Exist {
			m.Add(d.Key)
		}
		if m.IsExist(d.Key) != d.Exist {
			t.Fatalf("%s should be %t", d.Key, d.Exist)
		}
	}

}
