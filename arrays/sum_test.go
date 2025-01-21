package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Adds numbers from a collection of any size.", func(t *testing.T) {
		numbers := []int{10, 10, 10, 5} // Slice is a variable length array
		got := Sum(numbers)
		want := 35

		if got != want {
			t.Errorf("Got '%d', want '%d'", got, want)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 1, 1}, []int{2, 2, 2})
	want := []int{3, 6}

	if !slices.Equal(got, want) {
		t.Errorf("Got '%v', want '%v'", got, want)
	}

}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !slices.Equal(got, want) {
			t.Errorf("Got '%v', want '%v'", got, want)
		}
	}

	t.Run("Returns sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 6}, []int{4, 10, 10}, []int{2})
		want := []int{6, 20, 0}
		checkSums(t, got, want)
	})

	t.Run("Safely sums empty array", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		checkSums(t, got, want)
	})
}
