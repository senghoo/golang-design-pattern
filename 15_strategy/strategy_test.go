package strategy

import "testing"

func TestPayByCash(t *testing.T) {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash
}

func TestPayByBank(t *testing.T) {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
}
