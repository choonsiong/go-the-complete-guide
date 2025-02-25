package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		fmt.Println(name)
	}

	for range names {
		fmt.Println("looping...")
	}

	items := map[string]float64{
		"iPhone":  2999,
		"Mac Pro": 29999,
	}

	for item, price := range items {
		fmt.Println(item, price)
	}
}
