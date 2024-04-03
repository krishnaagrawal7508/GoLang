package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://jsonplaceholder.typicode.com:443/posts?userId=1&id=1"

func main() {
	fmt.Println("<---Welcome to Handling Urls Class--->")

	fmt.Println(myUrl)

	// parsing
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)   //protocol: https or http
	fmt.Println(result.Host)     //domain + port
	fmt.Println(result.Path)     //router
	fmt.Println(result.Port())   //port number
	fmt.Println(result.RawQuery) //query params

	//to get query param seperated
	qparams := result.Query()
	fmt.Printf("type of qparams: %T\n", qparams) //key pair (maps)
	fmt.Println(qparams)

	for key, val := range qparams {
		fmt.Printf("%v - %v\n", key, val)
	}

	//to create a url with given info
	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com:443",
		Path:     "/posts",
		RawQuery: "userId=1",
	}

	anotherUrl := partsofurl.String()
	fmt.Println(anotherUrl)

}
