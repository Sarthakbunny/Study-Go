package main

import (
	"bufio"
	"example/my-app/note"
	"fmt"
	"os"
	"strings"
)

func getNoteData() (string, string) {
	title := getDataWithSpaces("Enter the title of the note:")
	content := getDataWithSpaces("Enter the content of the note:")

	return title, content
}

func notesAppOperation() {
	title, content := getNoteData()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	note.Display()
	err = note.WriteNoteToFile()
	if err != nil {
		fmt.Println("There are some issues with saving note", err.Error())
		return
	}
	fmt.Println("Note written successfully")
}

func getDataWithSpaces(promptText string) string {
	fmt.Printf("%s ", promptText)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
