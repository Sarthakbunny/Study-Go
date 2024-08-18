package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFileOperation(fileName string) (float64, error) {
	data, error := os.ReadFile(fileName)

	if error != nil {
		errorString := fmt.Sprintln(`There was an error reading the file, please try again`)
		return 1000.0, errors.New(errorString)
	}
	valueText := string(data)
	value, error := strconv.ParseFloat(valueText, 64)

	if error != nil {
		errorString := fmt.Sprintln(`There was an error parsing the balance, please try again`)
		return 1000.0, errors.New(errorString)
	}

	return value, nil
}

func WriteFloatToFileOperation(fileName string, value float64) {
	balanceString := fmt.Sprintln(value)
	os.WriteFile(fileName, []byte(balanceString), 0644)
}

func WriteFloatsToFileOperation(fileName string, value1, value2, value3 float64) {
	balanceString := fmt.Sprintf("%0.2f\n%0.2f\n%0.2f", value1, value2, value3)
	os.WriteFile(fileName, []byte(balanceString), 0644)
}
