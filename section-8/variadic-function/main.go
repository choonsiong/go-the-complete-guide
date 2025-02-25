package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...))
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
