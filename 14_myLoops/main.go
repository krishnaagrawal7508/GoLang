package main

import (
	"fmt"
)

func main() {
	fmt.Println("<-----Welcome to Loop class----->")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	fmt.Println("<-----First type of Loop----->")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("<-----Second type of Loop----->")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("<-----Third type of Loop----->")
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	fmt.Println("<-----Fourth type of Loop----->")
	rougueValue := 1
	for rougueValue <= 5 {
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	fmt.Println("<-----Break----->")
	rougueValue = 1
	for rougueValue <= 5 {
		if rougueValue == 3 {
			break
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	fmt.Println("<-----Continue----->")
	rougueValue = 1
	for rougueValue <= 5 {
		if rougueValue == 3 {
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	fmt.Println("<-----Goto----->")
	rougueValue = 1
	for rougueValue <= 5 {
		if rougueValue == 3 {
			goto lco
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("jumping to lco from goto")

}
