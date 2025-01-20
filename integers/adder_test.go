package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Received '%d', expected '%d'", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(3, 7)
	fmt.Println(sum)
	// Output: 10
}

func ExampleAdd_second() {
	sum := Add(10, 10)
	subtracted := subtract(sum, 10)
	fmt.Println(subtracted)
	// Output: 10
}
