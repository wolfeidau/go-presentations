package main

import "testing"

func DoSomething() (int, error) {
	return 123, nil
}

func TestSomething(t *testing.T) {
	value, err := DoSomething()
	if err != nil {
		t.Error(err)
	}
	if value != 123 {
		t.Fatalf("expected 123, got: %d", value)
	}
}
