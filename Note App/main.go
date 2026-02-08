package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todoInstance, err := todo.New(todoText)
	if err != nil {
		fmt.Print(err)
		return
	}

	noteInstance, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	todoInstance.Display()
	err = todoInstance.Save()

	if err != nil {
		fmt.Println("Saving todo failed")
		return
	}

	fmt.Println("Saving todo succeeded")

	noteInstance.Display()
	err = noteInstance.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return
	}

	fmt.Println("Saving note succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf(" %v", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
