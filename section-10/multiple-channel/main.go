package main

import (
	"fmt"
	"time"
)

func main() {
	//dones := make([]chan bool, 4)

	//dones[0] = make(chan bool)
	//go slowGreet("How ... are ... you ...", dones[0])

	//dones[1] = make(chan bool)
	//go greet("Hello, World!", dones[1])

	//dones[2] = make(chan bool)
	//go greet("How are you?", dones[2])

	//dones[3] = make(chan bool)
	//go greet("Good bye!", dones[3])

	//for _, ch := range dones {
	//	<-ch
	//}

	done := make(chan bool)

	go slowGreet("How ... are ... you ...", done)
	go greet("Hello, World!", done)
	go greet("How are you?", done)
	go greet("Good bye!", done)

	for val := range done {
		fmt.Println(val)
	}
}

func greet(phrase string, ch chan bool) {
	fmt.Println(phrase)
	ch <- true
}

func slowGreet(phrase string, ch chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	ch <- true
	close(ch)
}
