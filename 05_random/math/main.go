package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to math, random, crypto")

	//random number geneartion through math algo
	fmt.Println("random number geneartion through math algo")
	//argument that is passed to (rand.seed || rand.NewSource) determines the value of rand.intn
	//to randomise the argument we are using the time whcih changes simultaneously
	//UnixNano returns t as a Unix time, the number of nanoseconds elapsed since January 1, 1970 UTC
	//rand.intn(x) return an integer value from [0,x)
	rand.NewSource(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1) // generates [1,5]

}
