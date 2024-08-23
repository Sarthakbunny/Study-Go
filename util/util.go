package util

import "fmt"

func GetInput(statement string) float64 {
	fmt.Println(statement)
	var inputValue float64
	fmt.Scanln(&inputValue)

	return inputValue
}
