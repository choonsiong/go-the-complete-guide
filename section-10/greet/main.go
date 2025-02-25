package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("Hello, World!")
	go greet("How are you?")
	go slowGreet("How ... are ... you ...")
	go greet("Good bye!")

	time.Sleep(5 * time.Second)
}

func greet(phrase string) {
	fmt.Println(phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
}
