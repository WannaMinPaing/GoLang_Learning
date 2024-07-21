package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName  string
	middleName string
	lastName   string
	createdAt  time.Time
}

func main() {
	userFirstName := getUserData("First Name")
	userMiddleName := getUserData("Middle Name")
	userLastName := getUserData("Last Name")

	var appUser user

	appUser = user{
		firstName:  userFirstName,
		middleName: userMiddleName,
		lastName:   userLastName,
		createdAt:  time.Now(),
	}

	// getUserDetails(appUser)
	appUser.getUserDetailsTwo()
	appUser.clearName()
	appUser.getUserDetailsTwo()

}

func getUserDetails(u user) {
	fmt.Println(u.firstName, u.middleName, u.lastName)
}

func (u *user) getUserDetailsTwo() {
	fmt.Println(u.firstName, u.middleName, u.lastName)
}

func (u *user) clearName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
