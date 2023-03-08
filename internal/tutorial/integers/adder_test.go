package integers

import (
	"fmt"
	"testing"
)

func Test_Adder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func Example_Add() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
	// ↑ this test won't be executed if you remove this
}
