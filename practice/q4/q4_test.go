package q4

import "testing"

func TestIncrement(t *testing.T) {
	x := 0
	Increment(&x)
	if x != 1 {
		t.Errorf("Increment(&x) = %d; want 1", x)
	}

	x = -1
	Increment(&x)
	if x != 0 {
		t.Errorf("Increment(&x) = %d; want 0", x)
	}

	x = 10
	Increment(&x)
	if x != 11 {
		t.Errorf("Increment(&x) = %d; want 11", x)
	}
}