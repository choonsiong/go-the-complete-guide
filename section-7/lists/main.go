package main

import "fmt"

func main() {
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr1)

	s1 := []string{"red", "blue", "green"}
	fmt.Println(s1)

	fmt.Println(arr1[:])
	fmt.Println(arr1[1:3])
	fmt.Println(arr1[:len(arr1)])

	s2 := []string{"red", "blue", "green"}
	s1 = append(s1, s2...)
	fmt.Println(s1)

	s3 := make([]int, 3000)
	fmt.Println(len(s3), cap(s3))

	s4 := []string{} // slice with 0 length and 0 capacity
	fmt.Println(s4 == nil)
	fmt.Println(len(s4), cap(s4))
	//s4[0] = "a"      // not ok
	s4 = append(s4, "a")
	fmt.Println(s4)

	var s5 []string // s5 is nil
	fmt.Println(s5 == nil)
	fmt.Println(len(s5), cap(s5)) // 0, 0
	//s5[0] = "a"                   // nok ok
	s5 = append(s5, "b")
	fmt.Println(s5)
}
