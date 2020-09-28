package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test Deposit", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Test Withdrawal", func (t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func (t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(200))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance (t *testing.T, w Wallet, expected Bitcoin) {
	t.Helper()

	got := w.Balance()

	if got != expected {
		t.Errorf("Expected %s, Got %s", expected, got)
	}
}

func assertNoError (t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Unexpected Error")
	}
}

func assertError (t *testing.T, err error, expected error) {
	t.Helper()

	if err == nil {
		t.Fatal("Expected Error, received nil instead")
	}

	if err != expected {
		t.Errorf("Expected %q, got %q", expected, err.Error())
	}
}
