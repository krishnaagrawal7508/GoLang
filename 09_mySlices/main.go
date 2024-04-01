package main

import (
	"fmt"
	"sort"
)

//Array and slice are two important data structures in Go that are used to store and manipulate sequences of values.
//Differences:
//An array is declared using the [size]type syntax, while a slice is displayed using the []type syntax.
//An array is a fixed-size data structure that stores a series of values of the same type, while a slice is a variable-size data structure that references a portion of an array.
//An array is a value type, while a slice is a reference type. Finally, an array is passed by value to a function, while a slice is passed by reference.

func main() {
	fmt.Println("<-----Welcome to slices class----->")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana", "Kiwi")
	fmt.Println("fruit list:", fruitList)

	//lower limit is inclusive and upper limit is exclusive
	//0 based indexing in both slices and arrays
	fruitList = append(fruitList[1:])
	fmt.Println("fruit list:", fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println("fruit list:", fruitList)

	fruitList = append(fruitList[:2])
	fmt.Println("fruit list:", fruitList)

	//append removes the data only
	fmt.Println("fruit list:", fruitList[0])

	//another way to make a slice
	fmt.Println("<-----Some functions on Slices----->")
	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777 => will give error, size is 4

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	//sorting
	sort.Ints((highScore))
	fmt.Println(highScore)
	fmt.Println((sort.IntsAreSorted(highScore)))

	//Remove a value from slices based on index
	fmt.Println("<-----Remove a value from slices based on index----->")

	var courses = []string{"python", "reactjs", "swift", "ruby", "java"}
	fmt.Println("Course list:", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Course list:", courses)

}
