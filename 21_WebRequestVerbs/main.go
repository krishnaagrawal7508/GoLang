package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("<----welcome to Web Request Class---->")

	fmt.Println("<-------GET------->")
	PerformGetRequests()

	fmt.Println("<----POST JSON DATA---->")
	PerformPost_JSON_Requests()

	fmt.Println("<----POST FORM DATA---->")
	PerformPost_FORM_Requests()

}

// GET requests
func PerformGetRequests() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// POST request as a json data
func PerformPost_JSON_Requests() {
	const myUrl = "http://localhost:3000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's Go With GoLang",
			"price" : 0,
			"platform" : "LearnCodeOnline.in"
		}	
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// POST request as a form data
func PerformPost_FORM_Requests() {
	const myUrl = "http://localhost:3000/postform"

	//form data
	data := url.Values{}
	data.Add("firstname", "krishna")
	data.Add("lastname", "agrawal")
	data.Add("email", "krishna@go.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
