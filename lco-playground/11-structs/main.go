package main

import "fmt"

func main()  {
	ronith := Person{"Ronith", 4, "ronith.r@yopmail.com"}
	fmt.Println(ronith)
	fmt.Printf("Individual values: %+v\n", ronith)
	fmt.Printf("Name: %v, Age: %v\n", ronith.Name, ronith.Age)
}

type Person struct {
	Name  string
	Age   int
	Email string
}