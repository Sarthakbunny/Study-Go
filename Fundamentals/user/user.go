package user

import (
	"errors"
	"fmt"
	"time"
)

type Str string

func (s *Str) Log() {
	fmt.Println(s)
}

type User struct {
	firstName string
	lastName  string
	birthdate string
	CreatedAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "-----",
			CreatedAt: time.Now(),
		},
	}
}

func (u *User) OutputDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName and lastName and birthdate are required")
	}
	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}
