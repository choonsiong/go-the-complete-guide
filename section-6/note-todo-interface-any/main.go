package main

import (
	"bufio"
	"example.com/note-todo-interface-any/note"
	"example.com/note-todo-interface-any/todo"
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

	printAnything(newNote)
	printAnything(newTodo)
}

func printAnything(value any) {
	//fmt.Println(value)
	printAnythingCondition(value)
}

func printAnythingCondition(value any) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	default:
		fmt.Println("Unknown type")
	}

	val, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", val)
	}

	val2, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", val2)
	}

	val3, ok := value.(string)
	if ok {
		fmt.Println("String: ", val3)
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
	err := saveData(o)
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
