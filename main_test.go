package main

import "testing"

func TestDouble(t *testing.T) {
	if got, want := double(2), 4; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
