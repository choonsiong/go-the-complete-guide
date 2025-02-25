package main

import (
	"bufio"
	"example.com/note-todo/note"
	"example.com/note-todo/todo"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	task := getTodoData()
	newTodo, err := todo.New(task)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()
	err = newNote.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	newTodo.Display()
	err = newTodo.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getTodoData() string {
	return getUserInput("Task: ")
}

func getUserInput(prompt string) (value string) {
	fmt.Print(prompt)
	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(value)
}
