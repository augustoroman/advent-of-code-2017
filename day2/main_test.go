package main

import (
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	testcases := []struct {
		str string
		sum int
	}{
		{"5 1 9 5\n7 5 3\n2 4 6 8\n", 18},
	}
	for _, test := range testcases {
		sum := sum(strings.NewReader(test.str))
		if sum != test.sum {
			t.Errorf("Expected sum of %d, got %d\nInput:\n%s",
				test.sum, sum, test.str)
		}
	}
}
