package money

type Money struct {
	currency string
	amount   int
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

func (augend *Money) plus(added Money) Expression {
	return Sum{*augend, added}
}
