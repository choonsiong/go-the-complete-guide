package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://www.google.com",
		"yahoo":  "https://www.yahoo.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])

	websites["tecnotree"] = "https://www.tecnotree.com"
	fmt.Println(websites)

	delete(websites, "tecnotree")
	fmt.Println(websites)

	var students map[string]int // nil map
	fmt.Println(students == nil)
	//students["alice"] = 10 // not ok

	colors := map[string]string{} // empty map
	fmt.Println(colors == nil)
	colors["red"] = "#ff0000" // ok
	fmt.Println(colors)

	names := make(map[string]string, 10)
	fmt.Println(names == nil)
}
