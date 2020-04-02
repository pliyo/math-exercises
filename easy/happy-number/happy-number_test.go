package main

import (
	"fmt"
	"testing"
)

func TestIsHappyNumber(t *testing.T) {
	testValues := []struct {
		inputValue    int
		expectedValue bool
	}{
		{1, true},
		{10, true},
		{13, true},
		{19, true},
		{44, true},
		{68, true},
		{91, true},
		{94, true},
		{2, false},
		{3, false},
		{4, false},
		{22, false},
		{24, false},
		{61, false},
	}

	for i, test := range testValues {
		fmt.Println("Testing with", test.inputValue)
		actualResult := IsHappyNumber(test.inputValue)

		if actualResult != test.expectedValue {
			t.Errorf("actualResult was incorrect got %t for %d when it should be %t (iteration: %d)", actualResult, test.inputValue, test.expectedValue, i+1)
		}
	}
}
