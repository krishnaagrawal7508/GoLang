package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	fmt.Println("<-----------LOCAL TIME--------->")
	//to get present local time
	presentTime := time.Now()
	fmt.Println(presentTime)
	//1)
	fmt.Println(presentTime.Format("01-02-2006"))
	//2)
	fmt.Println(presentTime.Format("15:04:05"))
	//3)
	fmt.Println(presentTime.Format("Monday"))
	//4)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	//5)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println("<-----------DESIRED TIME--------->")
	//to create a desired time
	createdDate := time.Date(2020, time.December, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006"))
	//2)
	fmt.Println(createdDate.Format("15:04:05"))
	//3)
	fmt.Println(createdDate.Format("Monday"))
	//4)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
	//5)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	// go env : command is used to get everything about the golang installed
	// there is GOOS also there which shows which os you are using
	// - for macos GOOS = darwin
	// - for windows GOOS = windows
	// - for linux GOOS = linux

	//to build a executable file command is - go build 
	//executable file is different for diff operating systems
	// for any desired os we can setup the GOOS
	// command: GOOS='darwin' go build
	// command: GOOS='windows' go build
	// command: GOOS='linux' go build

}
