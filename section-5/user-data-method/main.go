package main

import "fmt"

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthday := getUserData("Enter your birthday (MM/DD/YYYY): ")

	foobar, err := newUser(firstName, lastName, birthday)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(foobar.description())

	foobar.setBirthday("12/31/1990")
	fmt.Println(foobar.description())
}

func getUserData(prompt string) (value string) {
	fmt.Print(prompt)
	//_, _ = fmt.Scan(&value)
	_, _ = fmt.Scanln(&value)
	return
}
