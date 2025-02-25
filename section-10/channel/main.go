package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go slowGreet("How ... are ... you ...", done)
	greet("Hello, World!")
	greet("How are you?")
	greet("Good bye!")

	<-done // wait channel to emit data
}

func greet(phrase string) {
	fmt.Println(phrase)
}

func slowGreet(phrase string, ch chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	ch <- true
}
