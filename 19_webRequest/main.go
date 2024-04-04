package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://txti.es/"

func main() {
	fmt.Println("<---Welcome to Web Requests Class--->")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type %T\n\n\n", response)

	//caller's responsibility to close connection
	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// databyte is in byte format
	fmt.Println(string(databyte))

}
