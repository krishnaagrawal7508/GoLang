package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Println("Please rate our pizza from 1 to 5")

	//comma okay syntax
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating: ")

	// comma ok || error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input);
	fmt.Printf("Type of this rating is %T", input);
}
