package main

// using gorilla-mux for creating api or routers
// go get: added github.com/gorilla/mux v1.8.1

import (
	"fmt"
	"net/http"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename`
	CoursePrice int     `json:"price"`
	Author      *Author `json,"author"`
}

type Author struct {
	FullName string `json: fullname`
	Website  string `json: website`
}

// database as a slice
var courses []Course

// middleware, helper - file
func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	fmt.Println("<----welcome to building Api class---->")

}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Learn CodeOnLine</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
