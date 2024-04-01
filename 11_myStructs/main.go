package main

import (
	"fmt"
)

func main() {
	fmt.Println("<-----Welcome to Structs class----->")

	//no inheritance in golang, no super, parent, child

	Krishna := User{"krishna", "k@a.com", true, 18}
	fmt.Println(Krishna)
	fmt.Printf("Krishna details are: %v\n", Krishna)
	fmt.Printf("Email is %v", Krishna.Email)

}

// All elements and structs should start with a capital letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
