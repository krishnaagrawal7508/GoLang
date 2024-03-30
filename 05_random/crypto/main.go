package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	fmt.Println("Welcome to math, random, crypto")

	//random number geneartion through cryptography
	fmt.Println("random number geneartion through crypto algo")
	//big.NewInt(x) return an integer value from [0,x)
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum) // generates [0,5)

}
