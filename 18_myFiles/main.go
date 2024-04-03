package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("<---Welcome to Files Class--->")

	content := "This need to go in a file - LearnCode"

	file, err := os.Create("./18_myFiles/myfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./18_myFiles/myfile.txt")

}

func readFile(filename string) {

	//any file which is read is not in string format it is in byte format
	contentByte, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Printf("RAW data: \n")
	fmt.Println(contentByte)

	fmt.Printf("TEXT data: \n")
	fmt.Println(string(contentByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
		//used to exit the code and just show the error
	}
}
