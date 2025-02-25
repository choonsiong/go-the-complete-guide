package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	age := 42

	fmt.Println("original age:", age)
	fmt.Println("new age:", modifyAge(age))

	fmt.Println("original age:", age)
	modifyAgeInPlace(&age)
	fmt.Println("original age:", age)
}

func modifyAge(age int) int {
	return age + randomdata.Number(20)
}

func modifyAgeInPlace(age *int) {
	*age += randomdata.Number(20)
}
