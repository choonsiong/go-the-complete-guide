package main

import (
	"example.com/user-data-package/user"
	"fmt"
)

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthday := getUserData("Enter your birthday (MM/DD/YYYY): ")

	foobar, err := user.New(firstName, lastName, birthday)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(foobar.Description())

	foobar.SetBirthday("12/31/1990")
	fmt.Println(foobar.Description())
}

func getUserData(prompt string) (value string) {
	fmt.Print(prompt)
	//_, _ = fmt.Scan(&value)
	_, _ = fmt.Scanln(&value)
	return
}
