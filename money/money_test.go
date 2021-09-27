package money

import "testing"

func TestMoney(t *testing.T) {
	five := Dollar{5}
	five.time(2)
	if five.amount != 10 {
		t.Errorf("expected 10, but answer is %d", five.amount)
	}
}
