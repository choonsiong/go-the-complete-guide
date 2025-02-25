package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println(doSomething(numbers, func(i int) int {
		return i * 2
	}))

	fmt.Println(doSomething(numbers, func(i int) int {
		return i * 3
	}))

	fmt.Println(doSomethingGeneric(numbers, func(i int) int {
		return i * 4
	}))

	fmt.Println(doSomethingGeneric([]string{"a", "b", "c"}, func(s string) string {
		return s + s
	}))

	fmt.Println(doSomethingAny(numbers, func(i int) int {
		return i * 4
	}))

	fmt.Println(doSomethingAny([]string{"a", "b", "c"}, func(s string) string {
		return s + s
	}))
}

func doSomething(nums []int, f func(int) int) (result []int) {
	for _, num := range nums {
		result = append(result, f(num))
	}
	return
}

func doSomethingGeneric[T int | float64 | string](arr []T, f func(T) T) (result []T) {
	for _, val := range arr {
		result = append(result, f(val))
	}
	return
}

func doSomethingAny[T any](arr []T, f func(T) T) (result []T) {
	for _, val := range arr {
		result = append(result, f(val))
	}
	return
}
