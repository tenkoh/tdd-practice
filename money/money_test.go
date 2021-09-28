package money

import (
	"testing"
)

func TestDollar_times(t *testing.T) {
	five := Dollar{5}
	product := five.times(2)
	if product.amount != 10 {
		t.Errorf("expected 10, but answer is %d", five.amount)
	}
	product = five.times(3)
	if product.amount != 15 {
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
