package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// type Stringer interface {
// 	String() string
// }

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func main() {
	wallet := Wallet{balance: Bitcoin(0)}
	fmt.Println("Initial balance: ", wallet.Balance())

	wallet.Deposit(Bitcoin(10))
	fmt.Println("After deposit balance: ", wallet.Balance())

	err := wallet.Withdraw(Bitcoin(100))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After overdraw balance: ", wallet.Balance())

	err = wallet.Withdraw(Bitcoin(5))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After withdraw balance: ", wallet.Balance())
}
