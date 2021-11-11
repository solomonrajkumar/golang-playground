package main

import "fmt"

const token = "Ron"

func main()  {
	var userName string = "Raj"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.12313234542342121132131313123
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit
	var website = "RAJ"
	fmt.Println(website)
	// this will fail -> website = true

	numberOfUsers := 100
	fmt.Println(numberOfUsers)

	fmt.Println(token)
	numberOfUsers = 200 // can reassign

}