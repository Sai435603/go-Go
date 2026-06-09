package main

import (
	"errors"
	"fmt"
)

// creating error
func CreateError(message string) error {
	return fmt.Errorf("An error occurred: %s", message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// custom error type
type InsufficientFundsError struct {
	Amount  float64
	Balance float64
}

func (e *InsufficientFundsError) Error() string {
	return "Insufficient funds: attempted to withdraw " + fmt.Sprintf("%.2f", e.Amount) + " with a balance of " + fmt.Sprintf("%.2f", e.Balance)
}

func Withdraw(amount float64, balance float64) error {

	if amount > balance {
		return &InsufficientFundsError{Amount: amount, Balance: balance}
	}
	return nil
}

func main() {
	err := Withdraw(150.00, 100.00)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Withdrawal successful")
	}

	err = CreateError("Hey this is error")
	if err != nil {
		fmt.Println(err)
	}
	_, err = Divide(2.3, 0)
	if err != nil {
		fmt.Println(err)
	}
}
