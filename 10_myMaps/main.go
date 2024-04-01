package main

import (
	"fmt"
)

func main() {
	fmt.Println("<-----Welcome to Maps class----->")

	Languages := make(map[string]string)

	Languages["JS"] = "Javascript"
	Languages["RB"] = "Ruby"
	Languages["PY"] = "Python"

	fmt.Println("List of all languages:", Languages)
	fmt.Println("JS shorts for:", Languages["JS"])

	fmt.Println("<-----Delete element in maps----->")
	delete(Languages, "RB")
	fmt.Println("List of all languages:", Languages)

	fmt.Println("<-----Print all elemts in loop----->")
	for key, value := range Languages {
		fmt.Printf("For key %v, value uis %v\n", key, value)
	}

}
