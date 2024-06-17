package pointers_and_errors


import "fmt"
import "errors"


var ErrInsufficientFunds = errors.New("can't withdraw, insufficient funds")


type Bitcoin int


func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}


type Wallet struct {
  Amount Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
  w.Amount += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
  if amt > w.Amount {
    return ErrInsufficientFunds
  }
  w.Amount -= amt
  return nil
}

func (w *Wallet) Balance() Bitcoin {
  return w.Amount
}
