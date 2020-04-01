package main

import (
	"fmt"
	"testing"
)

func TestFindSingleNumber(t *testing.T) {
	testValues := []struct {
		inputValues   []int
		expectedValue int
	}{
		{[]int{0, 1, 1, 1}, 0},
		{[]int{2, 2, 1, 2}, 1},
		{[]int{2, 2, 1, 2, 4, 1, 1, 2, 1}, 4},
	}

	for i, test := range testValues {
		fmt.Println("Testing with", test.inputValues)
		actualResult := SingleNumber(test.inputValues)

		if actualResult != test.expectedValue {
			t.Errorf("actualResult was incorrect got %d should be %d (iteration: %d)", actualResult, test.expectedValue, i+1)
		}
	}
}
