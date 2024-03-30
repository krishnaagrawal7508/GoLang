package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Println("Please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)

	// new variable numRating is assigned int of input
	// Parse float is used to convert to floating point 64 bit
	// Trimspace is used to trim "4\n" to "4"
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// to encounter an error 
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
