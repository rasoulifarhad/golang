package fib

import "testing"

var fiboTests = []struct {
	n        int // input
	expected int // expected resualt
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 14},
}

type fibTest struct {
	n        int
	expected int
}

var fibTests = []fibTest{
	{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fiboTests {
		actual := Fib(tt.n)

		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
