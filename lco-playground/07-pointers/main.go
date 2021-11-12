package main

import "fmt"

func main()  {
	//var ptr *int
	//fmt.Println("value of ptr is ", ptr)

	myNum := 10
	var memAddress = &myNum
	fmt.Println("Address ", memAddress)
	fmt.Println("Value ", *memAddress)

	*memAddress = 20
	fmt.Println("Address ", memAddress)
	fmt.Println("Value ", *memAddress)
}