package main_test

import (
	"testing"
)

var fibTests = []struct {
	a        int
	expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{20, 10946},
	{42, 433494437},
}

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func BenchmarkSimple(b *testing.B) {
	fibo(32)
}
