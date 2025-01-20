package iteration

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	t.Run("should repeat string 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Got '%q', want '%q'", repeated, expected)
		}
	})
	t.Run("should repeat 0 times if 0 provided", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("Got '%q', want '%q'", repeated, expected)
		}
	})
	t.Run("should repeat something different back", func(t *testing.T) {
		repeated := Repeat("BP!", 2)
		expected := "BP!BP!"

		if repeated != expected {
			t.Errorf("Got '%q', want '%q'", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("HULK SMASH", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("DISAPPEAR!", 3)
	fmt.Println(repeated)
	// Output: DISAPPEAR!DISAPPEAR!DISAPPEAR!
}
