package money

import (
	"testing"
)

func TestDollar_times(t *testing.T) {
	five := Dollar{5}
	expected := Dollar{10}
	if expected != five.times(2) {
		t.Errorf("expected 10, but answer is %d", five.amount)
	}

	expected = Dollar{15}
	if expected != five.times(3) {
		t.Errorf("expected 15, but answer is %d", five.amount)
	}
}

func TestDollar_equals(t *testing.T) {
	a := Dollar{5}
	b := Dollar{5}
	c := Dollar{6}
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
