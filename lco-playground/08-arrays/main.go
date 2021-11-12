package main

import "fmt"

func main()  {
	var vegetables [4]string
	vegetables[0] = "Potato"
	vegetables[1] = "Tomato"
	vegetables[3] = "Radish"
	fmt.Println(vegetables)
	fmt.Println(len(vegetables))
	fmt.Println(cap(vegetables))

	var phones = [5]string{"apple", "samsung", "nokia"}
	fmt.Println(phones)
	fmt.Println(len(phones))
	fmt.Println(cap(phones))
}