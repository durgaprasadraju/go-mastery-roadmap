package solutions

import "errors"

var ErrInsufficientFunds = errors.New("insufficient funds")

// BankAccount with balance.
type BankAccount struct{ balance int }

func (a *BankAccount) Deposit(amount int) { a.balance += amount }

func (a *BankAccount) Withdraw(amount int) error {
	if amount > a.balance {
		return ErrInsufficientFunds
	}
	a.balance -= amount
	return nil
}

func (a BankAccount) Balance() int { return a.balance }