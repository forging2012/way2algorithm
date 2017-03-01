package hashtable

import (
	"testing"
)

func TestHashtable(t *testing.T) {
	cases := []struct {
		key string
		val int
	}{
		{"a", 1},
		{"a", 2},
		{"b", 3},
		{"d", 3},
	}

	table := New()
	for _, c := range cases {
		table.Set(c.key, c.val)
		if val, ok := table.Get(c.key); !ok || val.(int) != c.val {
			t.FailNow()
		}
	}

	for _, c := range cases {
		table.Delete(c.key)
		if _, ok := table.Get(c.key); ok {
			t.FailNow()
		}
	}
}
