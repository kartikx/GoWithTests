package main

import "testing"

func TestSum(t *testing.T) {

	checkSum := func (t *testing.T, numbers []int, got, expected int) {
		t.Helper()

		if got != expected {
			t.Errorf("Expected %d, Got %d, %v", expected, got, numbers)
		}
	}

	t.Run("Collection of 3 Numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		checkSum(t, numbers, got, expected)
	})
}