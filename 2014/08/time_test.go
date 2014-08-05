package main

import (
	"testing"
	"time"
)

// START1 OMIT
type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

type fakeClock struct{}

func (fakeClock) Now() time.Time {
	t, _ := time.Parse("02/Jan/2006 15:04:05 -0700", "01/Aug/2014 15:04:05 +1000")
	return t
}

func TestNow(t *testing.T) {
	clk := fakeClock{}

	unix := clk.Now().Unix()

	if clk.Now().Unix() != 1406869445 {
		t.Errorf("actual=%d expected=%d\n", unix, 1406869445)
	}
}

// STOP1 OMIT
