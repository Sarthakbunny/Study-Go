package investmentcalculator

import (
	"example/my-app/Fundamentals/util"
	"fmt"
	"math"
)

const inflationRate = 2.5

func CalculateInvestmentReturn() {
	investedAmount := util.GetInput("Enter the investment amount: ")
	expectedReturnRate := util.GetInput("Enter the expected return rate: ")
	years := util.GetInput("Enter the years of investment: ")

	finalValue, futureRealValue := calculateReturnValues(investedAmount, expectedReturnRate, years)

	// fmt.Println("Future value for the investment:", finalValue)
	// fmt.Println("Future value for the investment (adjusted to inflation):", futureRealValue)
	// fmt.Printf("Future value for the investment: %0.1f\n Future value for the investment (adjusted to inflation): %0.1f", finalValue, futureRealValue)

	formattedReturnValue := fmt.Sprintf("Future value for the investment: %0.1f\n", finalValue)
	formattedAdjustedReturnValue := fmt.Sprintf("Future value for the investment (adjusted to inflation): %0.1f\n", futureRealValue)

	fmt.Print(formattedReturnValue, formattedAdjustedReturnValue)
}

func calculateReturnValues(investedAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investedAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
