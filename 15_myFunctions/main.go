package main

import "fmt"

func main() {
	fmt.Println("<----Welcome to function in golang---->")

	greeter()

	var result = sum(3, 5)
	fmt.Println("Sum is:", result)

	var proResult, myMessage = proSum(1, 2, 3, 4)
	fmt.Println("Sum is:", proResult)
	fmt.Println("Message is:", myMessage)

}

func greeter() {
	fmt.Println("Hello from golang !!")
}

func sum(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proSum(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi second return"
}
