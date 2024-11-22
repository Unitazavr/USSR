package main

import "testing"

func TestCountMiddleValue(t *testing.T) {
	type test struct {
		matrix   [][]int
		expected float64
	}
	tests := []test{
		{[][]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}}, 4},
		{[][]int{{0, 0, 0}, {5, 5, 5}}, 2.5},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1},
	}
	for _, r := range tests {
		got := CountMiddleValue(r.matrix)
		expected := r.expected
		if got != expected {
			t.Errorf("got: %f, expected: %f", got, expected)
		}
	}
}
