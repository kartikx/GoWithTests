package main

import "testing"

func TestHello(t *testing.T) {

	/**
	 * No Need to declare this function outside, or in a separate package,
	 * as it will only be used here. Go allows you to declare functions within
	 * other functions, and assign them to variables.
	 */
	assertCorrectMessage := func (t *testing.T, got, want string) {
		// ? This ensures that the error shows up in the Caller.
		t.Helper()
		
		if got != want {
			t.Errorf("Got %q Want %q", got, want)
		}
	}

	t.Run("Saying hello to People", func (t *testing.T) {
		got := Hello("Steven")
		want := "Hello Steven!"
		
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello World!' when empty string is supplied", func (t *testing.T) {
		got := Hello("")
		want := "Hello World!"
		
		assertCorrectMessage(t, got, want)
	})
}