package example

import "testing"

func TestSum(t *testing.T) {
	const a, b, out = 1, 1, 2

	if x := Sum(a, b); x != out {
		t.Errorf("Sum(%v, %v) = %v, want %v", a, b, out, out)
	}
}
