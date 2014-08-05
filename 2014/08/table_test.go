package main

import "testing"

func Add(a int, b int) int {
	return a + b
}

// START1 OMIT
var addTest = []struct {
	a   int
	b   int
	sum int
}{
	{
		a: 1, b: 2, sum: 3,
	},
	{
		a: 500, b: 2222, sum: 2722,
	},
}

func TestExample(t *testing.T) {
	for _, ts := range addTest {
		res := Add(ts.a, ts.b)
		if res != ts.sum {
			t.Errorf("actual=%d expected=%d\n", res, ts.sum)
		}
	}
}

// STOP1 OMIT
