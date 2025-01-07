package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("vexi")
	expected := "Hello vexi!"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
