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

func TestFranc_times(t *testing.T) {
	five := Franc{5}
	expected := Franc{10}
	if expected != five.times(2) {
		t.Errorf("expected 10, but answer is %d", five.amount)
	}

	expected = Franc{15}
	if expected != five.times(3) {
		t.Errorf("expected 15, but answer is %d", five.amount)
	}
}

func TestFranc_equals(t *testing.T) {
	a := Franc{5}
	b := Franc{5}
	c := Franc{6}
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
