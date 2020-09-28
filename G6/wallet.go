package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Insufficient Funds")

// Bitcoin type.
type Bitcoin int

// Wallet struct.
type Wallet struct {
	balance Bitcoin
}

// Deposit Add amount to balance.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance Get the current Balance.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw the balance.
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func main() {

}
