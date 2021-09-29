package money

type Expression interface {
}

type Sum struct {
	augend Money
	added  Money
}

func (sum *Sum) reduce(to string) Money {
	amount := sum.augend.amount + sum.added.amount
	return NewMoney(to, amount)
}
