package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/08-methods/exercises/solutions"
)

func TestBankAccount(t *testing.T) {
	a := &solutions.BankAccount{}
	a.Deposit(100)
	if err := a.Withdraw(30); err != nil || a.Balance() != 70 {
		t.Fatalf("balance=%d err=%v", a.Balance(), err)
	}
}
