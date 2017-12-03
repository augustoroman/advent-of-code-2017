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
		lines := readLines(strings.NewReader(test.str))
		sum := sum(lines)
		if sum != test.sum {
			t.Errorf("Expected sum of %d, got %d\nInput:\n%s",
				test.sum, sum, test.str)
		}
	}
}
func TestSum2(t *testing.T) {
	testcases := []struct {
		str string
		sum int
	}{
		{"5 9 2 8\n9 4 7 3\n3 8 6 5", 9},
	}
	for _, test := range testcases {
		lines := readLines(strings.NewReader(test.str))
		sum := sum2(lines)
		if sum != test.sum {
			t.Errorf("Expected sum of %d, got %d\nInput:\n%s",
				test.sum, sum, test.str)
		}
	}
}
