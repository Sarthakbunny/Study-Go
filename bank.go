package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func readFileOperation() (float64, error) {
	data, error := os.ReadFile(accountBalanceFileName)

	if error != nil {
		errorString := fmt.Sprintln(`There was an error reading the file, please try again`)
		return 1000.0, errors.New(errorString)
	}
	balanceText := string(data)
	balance, error := strconv.ParseFloat(balanceText, 64)

	if error != nil {
		errorString := fmt.Sprintln(`There was an error parsing the balance, please try again`)
		return 1000.0, errors.New(errorString)
	}

	return balance, nil
}

func writeToFileOperation(balance float64) {
	balanceString := fmt.Sprintln(balance)
	os.WriteFile(accountBalanceFileName, []byte(balanceString), 0644)
}

func bankOperations() {
	accountBalance, error := readFileOperation()
	if error != nil {
		fmt.Println(`ERROR:`)
		fmt.Println(error.Error())
		fmt.Println(`-----------`)
		return
		// panic(error.Error())
	}
	var choice int

	for choice != 4 {
		fmt.Println(`Welcome to the bank service
		What do you like to do?
		1. Check you bank balance
		2. Make a deposit to the account
		3. Withdraw some money
		4. Exit`)

		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("You have bank balance of %0.2f\n", accountBalance)
		case 2:
			fmt.Println("How much do you want to deposit")
			var amountToDeposit float64
			fmt.Scan(&amountToDeposit)

			if amountToDeposit < 0 {
				fmt.Println("The amount entered is negative")
				continue
			}

			accountBalance += amountToDeposit
			fmt.Printf("You have bank balance of %0.2f\n", accountBalance)
			writeToFileOperation(accountBalance)
		case 3:
			fmt.Println("How much do you want to withdraw")
			var amountToWidthraw float64
			fmt.Scan(&amountToWidthraw)
			if amountToWidthraw < 0 {
				fmt.Println("The amount entered is negative")
				continue
			}

			if amountToWidthraw > accountBalance {
				fmt.Printf("You cannot withdraw more than %0.2f\n", accountBalance)
				continue
			}

			accountBalance -= amountToWidthraw
			fmt.Printf("You have bank balance of %0.2f\n", accountBalance)
			writeToFileOperation(accountBalance)
		case 4:
			fmt.Println("You exited")
		default:
			return
		}
	}
}
