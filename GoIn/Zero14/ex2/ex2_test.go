package ex2

import "testing"

type Test struct {
	in  int
	out string
}

// Тестовые случаи
var tests = []Test{
	{-1, "negative"},
	{0, "zero"},
	{5, "short"},
	{9, "short"},
	{10, "long"},
	{99, "long"},
	{100, "very long"},
	{150, "very long"},
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Length(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
