package wallet

import (
	"errors"
	"fmt"
)

//ErrInsufficientFunds chao e qu qian cuo wu
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Bitcoin new type from int to describe bitcoind
type Bitcoin int

//Wallet type include balance
type Wallet struct {
	balance Bitcoin
}

//Stringer interface
type Stringer interface {
	String() string
}

//String return bitcoin string
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Deposit add balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Withdraw quqian
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil

}

//Balance return wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
