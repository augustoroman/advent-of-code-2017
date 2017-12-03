package main

import "testing"

func TestSummer(t *testing.T) {
	if sum := sumRepeatedDigits("1111"); sum != 4 {
		t.Fatalf("Wrong sum: %d", sum)
	}
	if sum := sumRepeatedDigits("1122"); sum != 3 {
		t.Fatalf("Wrong sum: %d", sum)
	}
	if sum := sumRepeatedDigits("1234"); sum != 0 {
		t.Fatalf("Wrong sum: %d", sum)
	}
	if sum := sumRepeatedDigits("91212129"); sum != 9 {
		t.Fatalf("Wrong sum: %d", sum)
	}
}
func TestSummer2(t *testing.T) {
	testcases := []struct {
		str string
		sum int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}
	for _, test := range testcases {
		sum := sumOffsetRepeatedDigits(test.str)
		if sum != test.sum {
			t.Fatalf("sum(%q) = %d, but should = %d", test.str, sum, test.sum)
		}
	}
}
