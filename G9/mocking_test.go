package main

import (
	"bytes"
	"reflect"
	"testing"
)

/*
	In the first test, we Check if the Output is correct, by passing
	in the Output stream as the Bytes Buffer.
	In the second test, we Check if the Order of the Operations is correct,
	by passing in the Spy as Writer as well as Sleeper.
*/
func TestCountdown(t *testing.T) {
	t.Run("Test Output", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Countdown(&buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if want != got {
			t.Errorf("Want %q, got %q", want, got)
		}
	})

	t.Run("Test Operations Order", func(t *testing.T) {
		cSpy := &CountdownOperationsSpy{}

		Countdown(cSpy, cSpy)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, cSpy.Calls) {
			t.Errorf("wanted calls %v, got %v", want, cSpy.Calls)
		}
	})
}
