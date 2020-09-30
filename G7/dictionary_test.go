package main

import (
	"fmt"
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("Word exists in Dictionary", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is a test"

		if err != nil {
			t.Fatal("Unexpected Error")
		}

		assertStrings(t, got, want)
	})

	t.Run("Word doesn't exist in Dictionary", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrWordDoesNotExist

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Add("test", "this is a test")

	got, err := dictionary.Search("test")
	want := "this is a test"

	assertNoError(t, err)
	assertStrings(t, got, want)
}

// Helper Functions.

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Unexpected Error, Aborting")
	}
}

/**
 * If we use this function, we don't even need to check if got == nil,
 * because we're not invoking the .Error() on it.
 */
func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Expected error %q, got %q", want, got)
	}
}
