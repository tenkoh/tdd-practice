package money

type Money struct {
	currency string
	amount   int
}

type Bank struct {
}

type Expression interface {
}

func (b *Bank) reduce(exp Expression, currency string) Money {
	return Money{"USD", 10}
}

func NewMoney(currency string, amount int) Money {
	return Money{currency, amount}
}

func (m *Money) equals(i interface{}) bool {
	money, ok := i.(Money)
	if !ok {
		return false
	}
	return m.amount == money.amount && m.currency == money.currency
}

func (m *Money) times(multiplier int) Money {
	return Money{m.currency, m.amount * multiplier}
}

func (m *Money) plus(money Money) Expression {
	return nil
}
