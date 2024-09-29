package main

import (
    "testing"
)

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {-1, -1, -2},
        {100, 200, 300},
    }

    for _, test := range tests {
        result := Sum(test.a, test.b)
        if result != test.expected {
            t.Errorf("Sum(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
        }
    }
}
