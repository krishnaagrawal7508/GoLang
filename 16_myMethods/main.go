package main

import (
	"fmt"
)

//Method is a function defined for a struct

func main() {
	fmt.Println("<-----Welcome to Methods class----->")

	Krishna := User{"krishna", "k@a.com", true, 18}
	fmt.Println(Krishna)
	fmt.Printf("Krishna details are: %v\n", Krishna)
	fmt.Printf("Email is %v\n\n", Krishna.Email)

	// calling a method
	// copy of the user is passed - call by value
	Krishna.GetStatus()
	Krishna.NewMail()

	// copy of the user is passed - call by value
	fmt.Printf("Email is %v\n\n", Krishna.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// defining a method
// copy of the user is passed
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "techdev@gmail.com"
	fmt.Println("New email of user: ", u.Email)
}
