package bank

import (
	"fmt"

	"example/my-app/Fundamentals/fileOps"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFileName = "Balance.txt"

func BankOperations() {
	accountBalance, error := fileOps.ReadFloatFileOperation(accountBalanceFileName)
	if error != nil {
		fmt.Println(`ERROR:`)
		fmt.Println(error.Error())
		fmt.Println(`-----------`)

		fmt.Println("Using the default balance")
		fileOps.WriteFloatToFileOperation(accountBalanceFileName, accountBalance)
		return
		// panic(error.Error())
	}
	var choice int

	for choice != 4 {
		initCallFromBank()

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
			fileOps.WriteFloatToFileOperation(accountBalanceFileName, accountBalance)
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
			fileOps.WriteFloatToFileOperation(accountBalanceFileName, accountBalance)
		case 4:
			fmt.Println("You exited")
		default:
			return
		}
	}
}

func initCallFromBank() {
	fmt.Println("Welcome to the bank service")
	fmt.Printf("Call us 24/7 @%s\n", randomdata.PhoneNumber())
	fmt.Println("What do you like to do?")
	fmt.Println("1. Check you bank balance")
	fmt.Println("2. Make a deposit to the account")
	fmt.Println("3. Withdraw some money")
	fmt.Println("4. Exit")
}
