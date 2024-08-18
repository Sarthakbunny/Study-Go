package main

import (
	"fmt"
	"os"
)

const expenseFileName = "ExpenseData.txt"

func writeToExpenseFile(ebt, eat, ratio float64) {
	resString := fmt.Sprintf("%0.2f \n%0.2f \n%0.2f", ebt, eat, ratio)
	os.WriteFile(expenseFileName, []byte(resString), 0644)
}

func calculateProfit() {
	revenue := getInput("Enter the revenue generated: ")
	expenses := getInput("Enter the expenses occurred: ")
	tax := getInput("Enter the tax rates calculated: ")

	if revenue <= 0 || expenses <= 0 || tax <= 0 {
		fmt.Println(`The input value must be greater than or equal to 0`)
		return
	}

	ebt, eat, ratio := calculateEBTs(expenses, revenue, tax)

	writeToExpenseFile(ebt, eat, ratio)
}

func getInput(statement string) float64 {
	fmt.Println(statement)
	var inputValue float64
	fmt.Scanln(&inputValue)

	return inputValue
}

func calculateEBTs(expenses, revenue, tax float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - tax/100)
	ratio := ebt / eat

	return ebt, eat, ratio

}
