package money

type Money struct {
	amount int
}

func (m *Money) equals(i interface{}) bool {
	money, ok := i.(Money)
	if !ok {
		return false
	}
	return m.amount == money.amount
}

func (m *Money) times(multiplier int) Money {
	return Money{m.amount * multiplier}
}

type Dollar struct {
	amount int
}
type Franc struct {
	amount int
}
