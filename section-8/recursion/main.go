package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial2(5))
}

func factorial(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

func factorial2(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
