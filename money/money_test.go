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

	if com := NewMoney("USD", 5); com != fiveDollar {
		t.Error("fail in equality check between dollars")
	}
	if com := NewMoney("USD", 6); com == fiveDollar {
		t.Error("fail in equality check between dollars")
	}
	if com := NewMoney("CHF", 5); com != fiveFranc {
		t.Error("fail in equality check between francs")
	}
	if com := NewMoney("CHF", 5); com == fiveDollar {
		t.Error("fail in equality check between dollar and franc")
	}
}
