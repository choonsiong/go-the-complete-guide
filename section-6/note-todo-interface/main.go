package main

import (
	"bufio"
	"example.com/note-todo-interface/note"
	"example.com/note-todo-interface/todo"
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

	err = displayData(newNote)
	if err != nil {
		return
	}

	err = displayData(newTodo)
	if err != nil {
		return
	}
}

func saveData(s Saver) error {
	err := s.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func displayData(o Outputable) error {
	o.Display()
	err := o.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
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
