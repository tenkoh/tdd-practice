package money

type Bank struct {
}

func (b *Bank) reduce(exp Expression, to string) Money {
	sum := exp.(Sum)
	return sum.reduce(to)
}
