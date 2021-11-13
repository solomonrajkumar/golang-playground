package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main()  {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Type : %T\n", response)

	// mandatory
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response:", string(databytes))
}