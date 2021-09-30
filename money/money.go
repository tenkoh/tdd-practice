package money

import "fmt"

type Money struct {
	currency string
	amount   int
}

func NewMoney(currency string, amount int) *Money {
	return &Money{currency, amount}
}

func (m *Money) Equals(i interface{}) bool {
	money, ok := i.(*Money)
	if !ok {
		return false
	}
	return m.amount == money.amount && m.currency == money.currency
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{m.currency, m.amount * multiplier}
}

func (augend *Money) Plus(added *Money) Expression {
	return &Sum{augend, added}
}

func (m *Money) Reduce(to string) *Money {
	return NewMoney(to, m.amount)
}

func (m *Money) ToString() string {
	return fmt.Sprintf("%d %s", m.amount, m.currency)
}
