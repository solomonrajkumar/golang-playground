package main

import "fmt"

func main()  {
	isLoggedIn := 10 > 4

	if isLoggedIn {
		fmt.Println("Logged in")
	} else {
		fmt.Println("Logged out")
	}

	if num := 4; num < 10 {
		fmt.Println("num < 10")
	} else {
		fmt.Println("num >= 10")
	}
}