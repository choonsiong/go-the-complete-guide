package main

import "fmt"

type stringMap map[string]string

func (s stringMap) Output() {
	fmt.Println(s)
}

func main() {
	students := make(stringMap, 10)

	students["alice"] = "A"
	students["bob"] = "B"
	students["charlie"] = "C"

	students.Output()
}
