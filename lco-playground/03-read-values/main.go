package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating: ")
	rating, _ := reader.ReadString('\n')
	fmt.Printf("Type is %T", rating)
}