package main

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser()

	if u == nil {
		t.Fatalf("Expected new user got nil")
	}
}
