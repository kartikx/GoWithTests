package greet

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(os.Stdout, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Expected: %q, got %q", want, got)
	}
}
