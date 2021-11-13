package main

import "fmt"

func main()  {
	fmt.Println(adder(1,2))
	fmt.Println(proAdder(1,2,3))
}

func adder(first int, second int) int {
	return first + second
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "hi"
}