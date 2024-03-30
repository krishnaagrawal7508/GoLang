package main

import "fmt"

// jwtToken := 30000 (this is not allowed outside function)

//global variable
const LoginToken string = "xyufgsuojbfavsdjfvbgnhfgdfs"

func main() {
	var username string = "krishna"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 234.124
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	//implicit type
	var website = "krishnaagrawal.com"
	fmt.Println(website + "\n")

	//no var style
	numberofUser := 300000
	fmt.Println(numberofUser)

	// global variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
