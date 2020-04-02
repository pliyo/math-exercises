package main

import (
	"testing"
)

func TestFindPoint(t *testing.T) {
	tables := []struct {
		px int32
		py int32
		qx int32
		qy int32
		rx int32
		ry int32
	}{
		{0, 0, 1, 1, 2, 2},
		{1, 1, 2, 2, 3, 3},
		{2, 2, 4, 4, 6, 6},
		{2, 5, 4, 4, 6, 3},
		{5, 5, 4, 4, 3, 3},
		{5, 2, 4, 4, 3, 6},
		{2, 3, 2, 3, 2, 3},
		{2, 3, 2, 4, 2, 5},
		{-2, 3, -4, 4, -6, 5},
	}

	for i, table := range tables {
		total := FindPoint(table.px, table.py, table.qx, table.qy)
		if total[0] != table.rx {
			t.Errorf("rx was incorrect got %d should be %d (iteration: %d)", total[0], table.rx, i+1)
		}
		if total[1] != table.ry {
			t.Errorf("ry was incorrect got %d should be %d (iteration: %d)", total[0], table.ry, i+1)
		}
	}
}
