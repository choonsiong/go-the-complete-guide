package main

import (
	"fmt"
	"time"
)

type UserData struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
	updatedAt time.Time
}

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthday := getUserData("Enter your birthday (MM/DD/YYYY): ")

	userData := UserData{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	outputUserData(userData)

	outputUserData(newUser(firstName, lastName, birthday))

	anotherUser := UserData{"alice", "smith", "10/10/1999", time.Now(), time.Now()}
	outputUserData(anotherUser)
}

func getUserData(prompt string) (value string) {
	fmt.Print(prompt)
	_, _ = fmt.Scan(&value)
	return
}

func outputUserData(userData UserData) {
	fmt.Println("First Name: ", userData.firstName)
	fmt.Println("Last Name: ", userData.lastName)
	fmt.Println("Birthday: ", userData.birthday)
	fmt.Println("Created: ", userData.createdAt)
	fmt.Println("Updated: ", userData.updatedAt)
}

func newUser(firstName string, lastName string, birthday string) UserData {
	return UserData{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}
