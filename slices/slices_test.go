package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 4}, []int{0, 9, 1})
		want := []int{6, 10}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
