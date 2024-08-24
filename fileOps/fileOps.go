package fileOps

import (
	"bufio"
	"encoding/json"
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

type FileManager struct {
	inputFileName  string
	outputFileName string
}

func New(inputFileName string, outputFileName string) FileManager {
	return FileManager{
		inputFileName:  inputFileName,
		outputFileName: outputFileName,
	}
}

func (fileManager FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fileManager.inputFileName)
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error while reading the file: ", scanner.Err())
		file.Close()
		return nil, err
	}

	return lines, nil
}

func (fileManager FileManager) WriteJSON(data any) error {
	file, err := os.Create(fileManager.outputFileName)

	if err != nil {
		fmt.Println("Error while creating the file: ", err)
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		fmt.Println("error while encoding: ", err)
		return err
	}
	file.Close()
	return nil
}
