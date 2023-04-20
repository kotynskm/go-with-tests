package wallet

import (
	"errors"
	"fmt"
)

// lowercase on a variable makes it private outside of this package
type Wallet struct {
	total Bitcoin
}

// we can create custom types in Go, to be more descriptive
// so now instead of int, we replace int with Bitcoin
type Bitcoin int

type Stringer interface {
	String() string
}

// we can implement methods on custom types, it's the same as it looks on a struct
// this method is defined in the fmt package and defines how your type is printed when using %s
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// whenever we use a receiver, we're making a copy of whatever we called the method from
// we can fix this with pointers. if we instead point to the wallet, then we can change the original values
// these are called struct pointers, and they are automatically dereferenced, so we don't need *w inside the method body
func (w *Wallet) Deposit(amount Bitcoin){
	fmt.Printf("address of total from deposit method is %v \n", &w.total)
	w.total += amount
}

// technically we don't need to point to the actual value here, a copy would be fine, but it's better to keep method receivers of the same type for consistency
func (w *Wallet) Balance() Bitcoin {
	return w.total
}

// lets declare err as a var so we can have single source of truth and not repeat the code throughout
// var allows this to be declared as global variable to the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.total {
		// errors.New creates a new error with message of your choosing
		// return errors.New("cannot withdraw, insufficient funds")
		return ErrInsufficientFunds
	}
	w.total -= amount
	return nil
}