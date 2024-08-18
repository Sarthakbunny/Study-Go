package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func calculateInvestmentReturn() {
	var investedAmount float64
	expectedReturnRate := 5.5
	var years float64

	outputStatment("Enter the investment amount: ")
	fmt.Scan(&investedAmount)

	outputStatment("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputStatment("Enter the years of investment: ")
	fmt.Scan(&years)

	finalValue, futureRealValue := calculateReturnValues(investedAmount, expectedReturnRate, years)

	// fmt.Println("Future value for the investment:", finalValue)
	// fmt.Println("Future value for the investment (adjusted to inflation):", futureRealValue)
	// fmt.Printf("Future value for the investment: %0.1f\n Future value for the investment (adjusted to inflation): %0.1f", finalValue, futureRealValue)

	formattedReturnValue := fmt.Sprintf("Future value for the investment: %0.1f\n", finalValue)
	formattedAdjustedReturnValue := fmt.Sprintf("Future value for the investment (adjusted to inflation): %0.1f\n", futureRealValue)

	fmt.Print(formattedReturnValue, formattedAdjustedReturnValue)
}

func outputStatment(statement string) {
	fmt.Print(statement)
}

func calculateReturnValues(investedAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investedAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
