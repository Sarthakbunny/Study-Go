package structure

import (
	"fmt"

	"example/my-app/Fundamentals/user"
)

func StructOperation() {
	var name user.Str = "MyName"
	name.Log()

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	userData, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	adminData := user.NewAdmin("admin@gmail.com", "admin")
	adminData.OutputDetails()

	userData.OutputDetails()

	userData.ClearUserName()
	userData.OutputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
