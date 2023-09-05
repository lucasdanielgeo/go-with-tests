package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// var assertEquality = func(t *testing.T, got, want Bitcoin) {
	// 	if got != want {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}
	// }

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(12))

		got := wallet.Balance()
		want := Bitcoin(12)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		wallet.Withdraw(Bitcoin(2))

		got := wallet.Balance()
		want := Bitcoin(8)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
