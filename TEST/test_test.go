package main

import (
	"fmt"
	"testing"
)

func MinInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestMinIntBasic(t *testing.T) {
	var a int = 2
	var b int = -2

	var c = MinInt(a, b)
	if c != b {
		t.Errorf("MinInt() failed, expected %d, got %d", a, c)
	}
}

func TestMinIntTable(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d,%d", test.a, test.b)
		t.Run(testname, func(t *testing.T) {
			ans := MinInt(test.a, test.b)
			if ans != test.want {
				t.Errorf("MinInt() failed, expected %d, got %d", test.want, ans)
			}
		})
	}
}

func BenchmarkMinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinInt(-2324, 1241232)
	}
}
