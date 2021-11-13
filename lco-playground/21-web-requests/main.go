package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	testGetRequestCall()
}

func testGetRequestCall() {
	const url = "http://localhost:8000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	// approach 1
	byteData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteData))

	// approach 2 - Revisit
	//var stringBuilder strings.Builder
	//byteData, _ = ioutil.ReadAll(response.Body)
	//byteCount, _ := stringBuilder.Write(byteData)
	//fmt.Println(byteCount)
	//stringBuilder.Write(byteData)
	//fmt.Println(stringBuilder.String())
}