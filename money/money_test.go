package money

import (
	"testing"
)

func TestMoneyString(t *testing.T) {
	five := NewMoney("USD", 5)
	if five.ToString() != "5 USD" {
		t.Error("ToString fails")
	}
}

func TestMultiple(t *testing.T) {
	five := NewMoney("USD", 5)

	expected := NewMoney("USD", 10)
	if calc := five.Times(2); expected.ToString() != calc.ToString() {
		t.Errorf("expected 10, but answer is %d", calc.amount)
	}

	expected = NewMoney("USD", 15)
	if calc := five.Times(3); expected.ToString() != calc.ToString() {
		t.Errorf("expected 15, but answer is %d", calc.amount)
	}
}

func TestEquality(t *testing.T) {
	fiveDollar := NewMoney("USD", 5)

	if fived := NewMoney("USD", 5); !fiveDollar.Equals(fived) {
		t.Error("fail in equality check between dollars")
	}
	if sixd := NewMoney("USD", 6); fiveDollar.Equals(sixd) {
		t.Error("fail in equality check between dollars")
	}
	if fivef := NewMoney("CHF", 5); fiveDollar.Equals(fivef) {
		t.Error("fail in equality check between dollar and franc")
	}
}

func TestSimpleAddition(t *testing.T) {
	fiveDollar := NewMoney("USD", 5)
	sum := fiveDollar.Plus(NewMoney("USD", 5))
	bank := new(Bank)
	exchanged := bank.Exchange(sum, "USD")
	if expected := NewMoney("USD", 10); !expected.Equals(exchanged) {
		t.Error("simple addition fails")
	}
}

func TestPlusReturnSum(t *testing.T) {
	fiveDollar := NewMoney("USD", 5)

	result := fiveDollar.Plus(NewMoney("USD", 5))
	sum, ok := result.(*Sum)
	if !ok {
		t.Error("could not cast Expression to *Sum")
	}

	if !sum.augend.Equals(fiveDollar) {
		t.Error("money.plus does not return sum struct")
	}
	if !sum.added.Equals(fiveDollar) {
		t.Error("money.plus does not return sum struct")
	}
}

func TestReduceSum(t *testing.T) {
	var sum Expression
	augend := NewMoney("USD", 4)
	added := NewMoney("USD", 3)
	sum = augend.Plus(added)

	bank := new(Bank)
	exchanged := bank.Exchange(sum, "USD")
	if expected := NewMoney("USD", 7); !expected.Equals(exchanged) {
		t.Error("reducing sum fails")
	}
}

func TestReduceMoney(t *testing.T) {
	five := NewMoney("USD", 5)
	bank := new(Bank)
	exchanged := bank.Exchange(five, "USD")
	if !exchanged.Equals(five) {
		t.Error("bank.reduce(money) error")
	}
}
