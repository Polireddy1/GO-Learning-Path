package main

import (
	"errors"
	"fmt"
)

// Define the Account struct
type Account struct {
	Owner   string
	Balance float64
}

// Method to deposit money into the account
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.Balance += amount
	return nil
}

// Method to withdraw money from the account
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

// Method to check the balance of the account
func (a *Account) CheckBalance() float64 {
	return a.Balance
}

// Method to transfer money between accounts
func (a *Account) Transfer(amount float64, target *Account) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	return target.Deposit(amount)
}

// Main function to demonstrate the methods
func main() {
	// Create two accounts
	account1 := &Account{Owner: "Alice", Balance: 1000}
	account2 := &Account{Owner: "Bob", Balance: 500}

	// Deposit money into Alice's account
	if err := account1.Deposit(200); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Deposited 200 to Alice's account. New balance: %.2f\n", account1.CheckBalance())
	}

	// Withdraw money from Bob's account
	if err := account2.Withdraw(100); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Withdrew 100 from Bob's account. New balance: %.2f\n", account2.CheckBalance())
	}

	// Transfer money from Alice to Bob
	if err := account1.Transfer(300, account2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Transferred 300 from Alice to Bob. Alice's new balance: %.2f, Bob's new balance: %.2f\n", account1.CheckBalance(), account2.CheckBalance())
	}

	// Attempt to withdraw more than the balance
	if err := account2.Withdraw(1000); err != nil {
		fmt.Println("Error:", err)
	}
}
