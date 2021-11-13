package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	content := "Contents to go inside test file"

	file, err := os.Create("test.txt")
	if err != nil {
		return 
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		return 
	}
	fmt.Println("Content Written and is of length ", length)
	defer file.Close()
	readFile()
}

func readFile() {
	file, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	fmt.Println("READING FILE")
	fmt.Println(string(file))
}