package main

import "fmt"

//in golang array is not used often
//there is another data type which is used instead

func main() {
	fmt.Println("Welcome to Array class")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomata"
	//fruitList[2] is not given any value, so it given " "
	fruitList[3] = "Peach"

	fmt.Println("Our fruit-list:", fruitList)
	fmt.Println("Length of fruit-list: ", len(fruitList))
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	
	var vegList = [5]string {"potato","beans","mushroom"}
	fmt.Println("Our veg-List:", vegList)
	fmt.Println("Length of veg-List: ", len(vegList))
	fmt.Printf("Type of fruitList is %T\n", vegList)
}
