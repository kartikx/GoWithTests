package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(t *testing.T, numbers []int, got, expected int) {
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

func TestSumAllTails(t *testing.T) {

	t.Run("Testing on Non empty Slices", func (t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, Got %v", expected, got)
		}
	})

	t.Run("Testing on Empty Slices", func (t *testing.T) {
		got := SumAllTails([]int{2}, []int{})
		expected := []int{0, 0}

		if !reflect.DeepEqual(got, expected) {	
			t.Errorf("Expected %v, Got %v", expected, got)
		}
	})
}

