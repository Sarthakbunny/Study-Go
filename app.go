package main

import (
	"example/my-app/fileOps"
	"example/my-app/price"
	"fmt"
)

// func main() {
// 	investmentcalculator.CalculateInvestmentReturn()
// 	profitcalculator.CalculateProfit()
// 	bank.BankOperations()
// 	structure.StructOperation()
// 	notes.NotesAppOperation()
// 	arrayPractice.ArrayPractice()
// 	mapPractice.MapPractice()
// 	funcPractice.FuncPractice()
// }

func main() {
	taxes := []float64{10, 30, 50}

	for _, tax := range taxes {
		fileManager := fileOps.New("prices.txt", fmt.Sprintf("converted_%.0f.json", tax))
		// cmdManager := cmdmanager.New()
		priceObject := price.NewPrice(tax, fileManager)
		priceObject.Process()
	}
}
