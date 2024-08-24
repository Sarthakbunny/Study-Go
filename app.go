package main

import (
	"example/my-app/arrayPractice"
	"example/my-app/bank"
	"example/my-app/fileOps"
	"example/my-app/funcPractice"
	investmentcalculator "example/my-app/investment_calculator"
	"example/my-app/mapPractice"
	"example/my-app/notes"
	"example/my-app/price"
	profitcalculator "example/my-app/profitCalculator"
	structure "example/my-app/struct"
	"fmt"
)

func main() {
	investmentcalculator.CalculateInvestmentReturn()
	profitcalculator.CalculateProfit()
	bank.BankOperations()
	structure.StructOperation()
	notes.NotesAppOperation()
	arrayPractice.ArrayPractice()
	mapPractice.MapPractice()
	funcPractice.FuncPractice()
	taxes := []float64{10, 30, 50}

	for _, tax := range taxes {
		fileManager := fileOps.New("prices.txt", fmt.Sprintf("converted_%.0f.json", tax))
		// cmdManager := cmdmanager.New()
		priceObject := price.NewPrice(tax, fileManager)
		priceObject.Process()
	}
}
