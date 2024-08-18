package main

import (
	"example/my-app/fileOps"
	"fmt"
)

const expenseFileName = "ExpenseData.txt"

func calculateProfit() {
	revenue := getInput("Enter the revenue generated: ")
	expenses := getInput("Enter the expenses occurred: ")
	tax := getInput("Enter the tax rates calculated: ")

	if revenue <= 0 || expenses <= 0 || tax <= 0 {
		fmt.Println(`The input value must be greater than or equal to 0`)
		return
	}

	ebt, eat, ratio := calculateEBTs(expenses, revenue, tax)

	fileOps.WriteFloatsToFileOperation(expenseFileName, ebt, eat, ratio)
}

func calculateEBTs(expenses, revenue, tax float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - tax/100)
	ratio := ebt / eat

	return ebt, eat, ratio

}
