package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointers study")

	// a pointer pointing towards a integer
	var ptr *int

	//nil by default	
	fmt.Println("value of pointer is ",ptr)

	// pointing our ptr to myNumber
	myNumber := 23
	ptr = &myNumber
	fmt.Println("value of ptr is ", ptr)
	fmt.Println("value of &ptr is ", &ptr)
	fmt.Println("value of *ptr is ", *ptr)

	//changing value of myNumber using ptr 
	*ptr = *ptr * 2;
	fmt.Println("value of *ptr is ", myNumber)


}
