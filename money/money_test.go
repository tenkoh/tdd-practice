package money

import "testing"

func TestMoney(t *testing.T) {
	five := Dollar{5}
	product := five.times(2)
	if product.amount != 10 {
		t.Errorf("expected 10, but answer is %d", five.amount)
	}
	product = five.times(3)
	if product.amount != 15 {
		t.Errorf("expected 15, but answer is %d", five.amount)
	}
}
