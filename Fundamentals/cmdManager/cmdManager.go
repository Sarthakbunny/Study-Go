package cmdmanager

import (
	"fmt"
)

type CmdManager struct {
}

func New() *CmdManager {
	return &CmdManager{}
}

func (cmdManager *CmdManager) ReadLines() ([]string, error) {
	input := -1
	lines := []string{}
	fmt.Println("Enter 0 to end the inputs")
	for input != 0 {
		var currInput string
		fmt.Println("Enter the price: ")
		fmt.Scanln(&currInput)
	}
	return lines, nil
}

func (cmdManager *CmdManager) WriteJSON(data any) error {
	fmt.Println(data)
	return nil
}
