package main

import "fmt"

func main()  {
	go greeter("raj")
	greeter("ron")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}