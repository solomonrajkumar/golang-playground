package main

import (
	"fmt"
	"sort"
)

func main()  {
	var fruits = []string{"Apple", "Banana", "Orange"}
	fmt.Printf("Type is %T\n", fruits)
	fruits = append(fruits, "Mango", "Strawberry")
	fmt.Println(fruits)
	//fruits = append(fruits[1:])
	//fmt.Println(fruits)
	//fruits = append(fruits[1:3])
	//fmt.Println(fruits)
	fruits = append(fruits[:3])
	fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 867
	fmt.Println(highScores)

	highScores = append(highScores, 333, 444, 555)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
}