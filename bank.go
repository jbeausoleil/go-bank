package main

import (
	"fmt"
	"github.com/jbeausoleil/go-bank/fileOperations"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileOperations.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("----------------------------------------")
		//panic("Cannot continue due to non-existing balance file.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount.  Deposit must be greater than 0.")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated.  Your account balance is now:", accountBalance)
			fileOperations.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount.  Deposit must be greater than 0.")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Invalid amount.  Please limit your withdraw to available funds.")
				continue
			}

			accountBalance -= withdrawlAmount
			fmt.Println("Balance updated.  Your account balance is now:", accountBalance)
			fileOperations.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye.")
			fmt.Println("Thanks for choosing our bank")
			return // will not allow code outside of loop to be executed - must be used for switch statements
			//break // exit program - will allow outside of loop to be executed
		}
		fmt.Println("----------------------------------------")
	}
}
