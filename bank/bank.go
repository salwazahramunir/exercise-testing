package bank

import (
	"errors"
)

type Account struct {
	AccountNumber string
	HolderName    string
	Balance       Balance
	CreatedYear   int
}

type Balance struct {
	Amount float64
}

func NewAccount(AccountNumber string, HolderName string, Amount float64, CreatedYear int) *Account {
	return &Account{
		AccountNumber: AccountNumber,
		HolderName:    HolderName,
		Balance: Balance{
			Amount: Amount,
		},
		CreatedYear: CreatedYear,
	}
}

func (a *Account) Deposit(amount float64) error {

	if amount <= 0 {
		return errors.New("the input you entered is invalid")
	}

	a.Balance.Amount += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the input you entered is invalid")
	}

	if amount > a.Balance.Amount {
		return errors.New("insufficient funds")
	}

	a.Balance.Amount -= amount

	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance.Amount
}

func (a *Account) Reset() {
	a.Balance.Amount = 50.00 // default saya set sebagai 50.00
}
