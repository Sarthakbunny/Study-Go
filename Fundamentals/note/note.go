package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid input")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Display() {
	fmt.Println("Note Title: ", note.Title)
	fmt.Println("Note Content: ", note.Content)
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	noteJson, err := json.Marshal(*note)
	if err != nil {
		return err
	}
	noteJsonString := string(noteJson)
	fmt.Println(noteJsonString)

	os.WriteFile(fileName, []byte(noteJsonString), 0644)
	return nil
}
