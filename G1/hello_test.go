package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Steven")
	want := "Hello Steven!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}