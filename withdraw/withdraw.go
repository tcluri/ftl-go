package withdraw

import (
	"errors"
)

func Withdraw(balance, amount int) (int, error) {
	if balance-amount < 0 {
		return 0, errors.New("cannot withdraw more than available balance")
	}
	return balance - amount, nil
}
