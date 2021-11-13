package main

import "fmt"

func main()  {
	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	//for d:= 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	//for i:= range days {
	//	fmt.Println(days[i])
	//}

	//for index, day := range days {
	//	fmt.Printf("Index: %v and Value: %v\n", index, day)
	//}

	invalidValue := 5

	for i:= 1; i < 10; i++ {
		if i == invalidValue {
			continue
		}
		fmt.Println(i)

	}

}