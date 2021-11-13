package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println("Start game")
	case 2:
		fmt.Println("2 spots")
		fallthrough
	case 3:
		fmt.Println("3 spots")
		fallthrough
	case 4:
		fmt.Println("4 spots")
	case 5:
		fmt.Println("5 spots")
	case 6:
		fmt.Println("6 spots play again")
		fallthrough
	default:
		fmt.Println("ROLL AGAIN")
	}
}