package main

import "testing"

func TestAdd(t *testing.T) {

	sum := add(10, 20)
	if sum != 30 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 30)
	}
	sub := subtract(40, 20)
	if sub != 20 {
		t.Errorf("Subtract was incorrect, got: %d, want: %d.", sub, 20)
	}
}
