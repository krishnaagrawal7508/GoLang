package main

import "fmt"

// stack all the defer statements then execute them in order at end of the code
// LIFO is followed

func main() {

	fmt.Println("<---Welcome to Defer Class--->")

	//last in first out

	defer fmt.Println("zero")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")

	def()

}

func def() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v\n", i)
	}
}
