package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // it shows coursename in json on key
	Price    int
	Platform string   `json:"website"`        // it shows website in json on key
	Password string   `json:"-"`              // to avoid usage of password as api
	Tags     []string `json:"tags,omitempty"` // omitempty is used to not show if field is empty
}

func main() {
	fmt.Println("<----welcome to Web Request Class---->")

	fmt.Println("<----Encode a JSON---->")
	EncodeJSON()

	fmt.Println("<----Decode a JSON---->")
	DecodeJSON()
}

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "dfc123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "rgc123", nil},
	}

	//pacakge this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	// finalJson, err := json.Marshal(lcoCourses) used to get raw json

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJSON() {

	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev","js"]	
		}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is NOT valid")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", lcoCourse)

	for k,v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
