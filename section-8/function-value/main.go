package main

import "fmt"

type intModifier func(int) int

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(processNumbers(nums, double))
	fmt.Println(processNumbers(nums, triple))
	fmt.Println(processNumbers2(nums, triple))
	fmt.Println(processNumbers2(nums, getIntFunction("4")))
}

func getIntFunction(option string) intModifier {
	switch option {
	case "2":
		return double
	case "3":
		return triple
	default:
		return single
	}
}

func processNumbers(numbers []int, f func(int) int) []int {
	var result []int
	for _, n := range numbers {
		result = append(result, f(n))
	}
	return result
}

func processNumbers2(numbers []int, f intModifier) []int {
	var result []int
	for _, n := range numbers {
		result = append(result, f(n))
	}
	return result
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func single(num int) int {
	return num
}
