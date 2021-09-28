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

type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}

func (d *Dollar) equals(i interface{}) bool {
	dollar, ok := i.(Dollar)
	if !ok {
		return false
	}
	return d.amount == dollar.amount
}

type Franc struct {
	amount int
}

func (f *Franc) times(multiplier int) Franc {
	return Franc{f.amount * multiplier}
}

func (f *Franc) equals(i interface{}) bool {
	franc, ok := i.(Franc)
	if !ok {
		return false
	}
	return f.amount == franc.amount
}
