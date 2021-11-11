package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main()  {
	fmt.Println("Welcome to Math")

	// Regular Random Function
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Int63n(5))

	// Crypto Random
	fmt.Println(rand.Int(rand.Reader, big.NewInt(6)))
}