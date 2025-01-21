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

func TestReturnSame(t *testing.T) {
	t.Run("Returns the same memory address", func(t *testing.T) {
		original := []int{1, 2, 3}
		got := ReturnSame(&original)
		if &original != got {
			t.Errorf("Got address '%p', want '%p'", got, &original)
		}
	})

	t.Run("Returns the same value", func(t *testing.T) {
		original := []int{1, 2, 3}
		got := ReturnSame(&original)
		if !slices.Equal(*got, original) {
			t.Errorf("Got '%v', want '%v'", *got, original)
		}
	})
}

func TestReturnCopy(t *testing.T) {
	original := []int{1, 2, 3}
	got := ReturnCopy(original)
	if !slices.Equal(original, got) {
		t.Errorf("Got address '%p', want '%p'", got, &original)
	}
	if &original == &got {
		t.Errorf("Got memory address '%p', which shouldn't match original '%p'", &got, &original)
	}
}
