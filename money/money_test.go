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

	if fived := NewMoney("USD", 5); !fiveDollar.equals(fived) {
		t.Error("fail in equality check between dollars")
	}
	if sixd := NewMoney("USD", 6); fiveDollar.equals(sixd) {
		t.Error("fail in equality check between dollars")
	}
	if fivef := NewMoney("CHF", 5); fiveDollar.equals(fivef) {
		t.Error("fail in equality check between dollar and franc")
	}
}

func TestSimpleAddition(t *testing.T) {
	fiveDollar := NewMoney("USD", 5)
	sum := fiveDollar.plus(NewMoney("USD", 5))
	bank := new(Bank)
	reduced := bank.reduce(sum, "USD")
	if expected := NewMoney("USD", 10); !expected.equals(reduced) {
		t.Error("simple addition fails")
	}
}

func TestPlusReturnSum(t *testing.T) {
	fiveDollar := NewMoney("USD", 5)
	var result Expression
	var sum Sum

	result = fiveDollar.plus(NewMoney("USD", 5))
	sum = result.(Sum)

	if sum.augend != fiveDollar {
		t.Error("money.plus does not return sum struct")
	}
	if sum.added != fiveDollar {
		t.Error("money.plus does not return sum struct")
	}
}

func TestReduceSum(t *testing.T) {
	var sum Expression
	augend := NewMoney("USD", 4)
	added := NewMoney("USD", 3)
	sum = augend.plus(added)

	bank := new(Bank)
	reduced := bank.reduce(sum, "USD")
	if expected := NewMoney("USD", 7); expected != reduced {
		t.Error("reducing sum fails")
	}
}

func TestReduceMoney(t *testing.T) {
	five := NewMoney("USD", 5)
	bank := new(Bank)
	reduced := bank.reduce(five, "USD")
	if reduced != five {
		t.Error("bank.reduce(money) error")
	}
}
