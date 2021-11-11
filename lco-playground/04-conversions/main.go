package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to Raing App.")
	fmt.Print("Please leave a rating between 1-5: ")
	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us with ", rating)
	parsedRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rating bumped to ", parsedRating + 1)
	}
}