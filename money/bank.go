package money

type Bank struct {
}

func (b *Bank) Exchange(exp Expression, to string) *Money {
	return exp.Reduce(to)
}
