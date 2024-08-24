package notes

import (
	"bufio"
	"example/my-app/note"
	"example/my-app/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func getNoteData() (string, string) {
	title := getDataWithSpaces("Enter the title of the note:")
	content := getDataWithSpaces("Enter the content of the note:")

	return title, content
}

func getTodoData() string {
	content := getDataWithSpaces("Enter the todo Note:")

	return content
}

func NotesAppOperation() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Anything")
	result := add(1, 2)
	fmt.Println(result)

	title, content := getNoteData()
	todoContent := getTodoData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	todo, err := todo.New(todoContent)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = outputData(note)
	if err != nil {
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func printSomething(value interface{}) { // value any
	intValue, ok := value.(int)
	if ok {
		fmt.Println("Integer value: ", intValue)
		return
	}
	floatValue, ok := value.(float64)
	if ok {
		fmt.Println("Float value: ", floatValue)
		return
	}
	stringValue, ok := value.(string)
	if ok {
		fmt.Println("String value: ", stringValue)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer value: ", value)
	// case float64:
	// 	fmt.Println("Float value: ", value)
	// case string:
	// 	fmt.Println("String value: ", value)
	// }
}

func outputData(data outputable) error {
	data.Display()
	return saveDataToFile(data)
}

func saveDataToFile(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("There are some issues with saving todo", err.Error())
		return err
	}
	fmt.Println("Todo written successfully")
	return nil
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
