package calculator

import "testing"

func TestAdd(t *testing.T) {
	if got, want := Add(1, 2), 3; got != want {
		t.Errorf("add function produced wrong result. expected: %d, got: %d", want, got)
	}
}

func TestSubtract(t *testing.T) {
	if got, want := Subtract(3, 2), 1; got != want {
		t.Errorf("Subtract function produced wrong result. expected: %d, got: %d", want, got)
	}
}

func TestMultiply(t *testing.T) {
	if got, want := Multiply(3, 2), 6; got != want {
		t.Errorf("Multiply function produced wrong result. expected: %d, got: %d", want, got)
	}
}
