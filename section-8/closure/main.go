package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	f1 := createFunc(2)
	f2 := createFunc(3)
	f3 := createFunc(4)

	fmt.Println(transform(numbers, f1))
	fmt.Println(transform(numbers, f2))
	fmt.Println(transform(numbers, f3))
}

func createFunc(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func transform(nums []int, f func(int) int) (result []int) {
	for _, n := range nums {
		result = append(result, f(n))
	}
	return
}
