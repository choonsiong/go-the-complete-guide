package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.23, 4.56))
	fmt.Println(add("hello", "world"))
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
