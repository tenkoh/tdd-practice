package money

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
