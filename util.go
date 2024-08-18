package main

import "fmt"

func getInput(statement string) float64 {
	fmt.Println(statement)
	var inputValue float64
	fmt.Scanln(&inputValue)

	return inputValue
}
