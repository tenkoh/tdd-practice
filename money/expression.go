package money

type Expression interface {
	Reduce(string) *Money
}

type Sum struct {
	augend *Money
	added  *Money
}

func (sum *Sum) Reduce(to string) *Money {
	amount := sum.augend.amount + sum.added.amount
	return NewMoney(to, amount)
}
