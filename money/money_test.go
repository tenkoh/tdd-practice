package money

import (
	"testing"
)

func TestMultiple(t *testing.T) {
	five := Money{5}
	expected := Money{10}
	if calc := five.times(2); expected != calc {
		t.Errorf("expected 10, but answer is %d", calc.amount)
	}

	expected = Money{15}
	if calc := five.times(3); expected != calc {
		t.Errorf("expected 15, but answer is %d", calc.amount)
	}
}

func TestEquality(t *testing.T) {
	a := Money{5}
	b := Money{5}
	c := Money{6}
	// check Golang spec.
	if a != b {
		t.Error("comparing structs fail")
	}

	// unit tests below
	if !a.equals(b) {
		t.Error("func: equals does not work")
	}
	if a.equals(c) {
		t.Error("func: equals does not work")
	}
}
