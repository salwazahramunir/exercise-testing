package main

import (
	"exercise_testing/bank"
	"fmt"
)

// soal exercise : https://hackmd.io/@enrinal/ByHXRMEMkg

func main() {
	account := bank.NewAccount("123456789", "Salwa Az-Zahra", 0.0, 2024)

	// Deposit money
	if err := account.Deposit(100.50); err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to withdraw money
	if err := account.Withdraw(50); err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to withdraw more than the balance
	if err := account.Withdraw(100); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Balance:", account.GetBalance())
}
