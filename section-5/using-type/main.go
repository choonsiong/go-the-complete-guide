package main

import (
	"fmt"
	"log"
)

type str string

func (s str) log() {
	log.Println(s)
}

func main() {
	var name str
	name = "John"
	fmt.Println(name)

	name.log()
}
