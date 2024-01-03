package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func readBalanceFromFile() float64 {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Printf("Error reading %v so using default: 0.00", accountBalanceFile)
		return 0.0
	}

	balanceText := string(data)
	balance, err2 := strconv.ParseFloat(balanceText, 64)
	if err2 != nil {
		fmt.Printf("Error parsing %v from the file %v so using default: 0.00", balanceText, accountBalanceFile)
		return 0.0
	}

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if err != nil {
		fmt.Printf("Error writing to  %v so nothing is saved", accountBalanceFile)
	}
}

func main() {
	var accountBalance = readBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")
	var showMenu = true
	for showMenu {
		choice := mainMenuSelect()

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %.2f\n", accountBalance)
		case 2:
			fmt.Print("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit must be greater than zero!")
			} else {
				accountBalance += depositAmount
				fmt.Printf("Balance after deposit: %.2f\n", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 3:
			fmt.Print("Withdrawal Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal must be greater than zero!")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Withdrawal is greater account balance!")
			} else {
				accountBalance -= withdrawAmount
				fmt.Printf("Balance after withdrawal: %.2f\n", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		default:
			showMenu = false
		}

		//if choice == 1 {
		//	fmt.Printf("Your balance is: %.2f\n", accountBalance)
		//} else if choice == 2 {
		//	fmt.Print("Deposit Amount: ")
		//	var depositAmount float64
		//	fmt.Scan(&depositAmount)
		//	if depositAmount <= 0 {
		//		fmt.Println("Deposit must be greater than zero!")
		//	} else {
		//		accountBalance += depositAmount
		//		fmt.Printf("Balance after deposit: %.2f\n", accountBalance)
		//	}
		//
		//} else if choice == 3 {
		//	fmt.Print("Withdrawal Amount: ")
		//	var withdrawAmount float64
		//	fmt.Scan(&withdrawAmount)
		//
		//	if withdrawAmount <= 0 {
		//		fmt.Println("Withdrawal must be greater than zero!")
		//	} else if withdrawAmount > accountBalance {
		//		fmt.Println("Withdrawal is greater account balance!")
		//	} else {
		//		accountBalance -= withdrawAmount
		//		fmt.Printf("Balance after withdrawal: %.2f\n", accountBalance)
		//	}
		//} else {
		//	break
		//}
	}

	fmt.Println("Thanks for choosing Go Bank!")
}

func mainMenuSelect() int {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Select Your Choice: ")
	fmt.Scan(&choice)
	return choice
}
