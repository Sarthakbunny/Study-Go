package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

const todoFileName = "todo.json"

func New(content string) (*Todo, error) {
	if content == "" {
		return nil, errors.New("invalid input")
	}
	return &Todo{
		Content: content,
	}, nil
}

func (note *Todo) Display() {
	fmt.Println("Note Content: ", note.Content)
}

func (note *Todo) Save() error {
	noteJson, err := json.Marshal(*note)
	if err != nil {
		return err
	}
	noteJsonString := string(noteJson)
	fmt.Println(noteJsonString)

	os.WriteFile(todoFileName, []byte(noteJsonString), 0644)
	return nil
}
