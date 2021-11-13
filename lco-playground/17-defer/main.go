package main

import "fmt"

func main()  {
	defer fmt.Println("Ron")
	fmt.Println("Raj")
	fmt.Println("Reva")
	deferTest()
}

func deferTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}