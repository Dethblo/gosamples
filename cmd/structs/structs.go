package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// a method added to the user struct
// since the function doesn't change the values in the struct, we don't need to use pointer in the Receiver
// but best practice is for all receiver to use the same for consistency
func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// a pointer must be used here in the Receiver to mutate the values otherwise a copy is passed
func (u *user) clearUserName() {
	u.lastName = ""
	u.firstName = ""
}

// a common design pattern in GO is to create a constructor (or factory) function to make instances
// so that additional validation logic may be performed.  Optionally: the return could be a pointer
// so that extra copies are not made.
func newUser(fn, ln, bd string) user {
	// return an instance by making a struct literal
	return user{
		firstName: fn,
		lastName:  ln,
		birthdate: bd,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	// create instance using the constructor (factory) function
	appUser = newUser(userFirstName, userLastName, userBirthdate)

	// create another instance with empty values
	// anotherUser := user {}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
