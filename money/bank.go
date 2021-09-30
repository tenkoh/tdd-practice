package money

type Bank struct {
}

func (b *Bank) reduce(exp Expression, to string) Money {
	if m, ok := exp.(Money); ok {
		return m
	}
	sum := exp.(Sum)
	return sum.reduce(to)
}
