package main

import (
	"example.com/gosamples/internal/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	// create instance using the constructor (factory) function (which returns a pointer)
	appUser = user.New(userFirstName, userLastName, userBirthdate)

	// create another instance with empty values
	// anotherUser := user {}

	// create instance using the constructor (factory) function (no pointer)
	admin := user.NewAdmin("ADMIN", "ISGOD", "----------", "admin@mycompany.com", "thepassword")
	fmt.Println("The Admin details:")
	admin.OutputUserDetails()

	fmt.Println("The Supplied User details:")
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
