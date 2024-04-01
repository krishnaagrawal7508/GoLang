package main

import (
	"fmt"
)

func main() {
	fmt.Println("<-----Welcome to If-Else----->")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "non-regular user"
	} else {
		result = "10 login user"
	}

	fmt.Println("User type:", result)

	// some mathemathical expression can be also done
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// variable can be defined on the go in golang
	if num:=3; num<10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

}
