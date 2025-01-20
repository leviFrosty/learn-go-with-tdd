package arrays

import "testing"

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
