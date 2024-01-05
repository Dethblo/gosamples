package user

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

// Admin an example of embedding a struct into another (thereby reusing the properties & methods of an existing struct
// in the new struct).
type Admin struct {
	Email    string
	Password string
	User     // embedding the User struct fields & methods (anonymously)
}

func NewAdmin(fn, ln, bd, email, password string) Admin {
	// return an instance by making a struct literal
	return Admin{
		Email:    email,
		Password: password,
		// another literal for the embedded struct
		User: User{
			FirstName: fn,
			LastName:  ln,
			Birthdate: bd,
			CreatedAt: time.Now(),
		},
	}
}

// OutputUserDetails a method added to the User struct
// since the function doesn't change the values in the struct, we don't need to use pointer in the Receiver
// but best practice is for all receiver to use the same for consistency
func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

// ClearUserName a pointer must be used here in the Receiver to mutate the values otherwise a copy is passed
func (u *User) ClearUserName() {
	u.LastName = ""
	u.FirstName = ""
}

// New a common design pattern in GO is to create a constructor (or factory) function to make instances
// so that additional validation logic may be performed.  This implementation illustrates returning a pointer
// so that extra copies are not made, but this is often not needed.
func New(fn, ln, bd string) *User {
	// return an instance by making a struct literal
	return &User{
		FirstName: fn,
		LastName:  ln,
		Birthdate: bd,
		CreatedAt: time.Now(),
	}
}
