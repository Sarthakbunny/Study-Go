package profitcalculator

import (
	"example/my-app/fileOps"
	"example/my-app/util"
	"fmt"
)

const expenseFileName = "ExpenseData.txt"

func CalculateProfit() {
	revenue := util.GetInput("Enter the revenue generated: ")
	expenses := util.GetInput("Enter the expenses occurred: ")
	tax := util.GetInput("Enter the tax rates calculated: ")

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
