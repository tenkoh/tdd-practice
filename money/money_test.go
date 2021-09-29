package money

import (
	"testing"
)

func TestMultiple(t *testing.T) {
	five := NewMoney("USD", 5)
	expected := NewMoney("USD", 10)
	if calc := five.times(2); expected != calc {
		t.Errorf("expected 10, but answer is %d", calc.amount)
	}

	expected = NewMoney("USD", 15)
	if calc := five.times(3); expected != calc {
		t.Errorf("expected 15, but answer is %d", calc.amount)
	}
}

func TestEquality(t *testing.T) {
	fiveDollar := NewMoney("USD", 5)
	fiveFranc := NewMoney("CHF", 5)

	if fived := NewMoney("USD", 5); !fiveDollar.equals(fived) {
		t.Error("fail in equality check between dollars")
	}
	if sixd := NewMoney("USD", 6); fiveDollar.equals(sixd) {
		t.Error("fail in equality check between dollars")
	}
	if fivef := NewMoney("CHF", 5); !fiveFranc.equals(fivef) {
		t.Error("fail in equality check between francs")
	}
	if fivef := NewMoney("CHF", 5); fiveDollar.equals(fivef) {
		t.Error("fail in equality check between dollar and franc")
	}
}
