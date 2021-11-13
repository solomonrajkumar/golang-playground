package main

import "fmt"

func main()  {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["TS"] = "Typescript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("Value for TS: ", languages["TS"])

	delete(languages, "PY")
	fmt.Println(languages)


	for key, value := range languages {
		fmt.Printf("%v  %v\n", key, value)
	}
}