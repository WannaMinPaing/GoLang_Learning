package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
	"example.com/main/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// todo.Display()
	// err = saveData(todo)

	err = outputData(todo)

	if err != nil {
		return
	}

	// userNote.Display()
	// err = saveData(userNote)
	outputData(userNote)

}

func printSomething(value interface{}) {
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {

	err := data.Save()

	if err != nil {
		fmt.Println("Saveing the todo failed.")
		return err
	}

	fmt.Println("Saving the todo succeeded!.")

	return nil
}

func getNoteData() (string, string) {

	title := getUserInput("Title : ")
	content := getUserInput("Content : ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

/// 93
