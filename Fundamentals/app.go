package main

import (
	"example/my-app/Fundamentals/arrayPractice"
	"example/my-app/Fundamentals/bank"
	"example/my-app/Fundamentals/funcPractice"
	investmentcalculator "example/my-app/Fundamentals/investment_calculator"
	"example/my-app/Fundamentals/mapPractice"
	"example/my-app/Fundamentals/notes"
	profitcalculator "example/my-app/Fundamentals/profitCalculator"
	structure "example/my-app/Fundamentals/struct"
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
