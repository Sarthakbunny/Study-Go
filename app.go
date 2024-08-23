package main

import (
	"example/my-app/arrayPractice"
	"example/my-app/bank"
	"example/my-app/funcPractice"
	investmentcalculator "example/my-app/investment_calculator"
	"example/my-app/mapPractice"
	"example/my-app/notes"
	profitcalculator "example/my-app/profitCalculator"
	structure "example/my-app/struct"
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
}
