package main

import (
	"fmt"
	"net/url"
)

const url1 = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=w32wd23d"

func main()  {
	// parsing url
	response, err := url.Parse(url1)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Scheme)
	fmt.Println(response.Host)
	fmt.Println(response.Path)
	fmt.Println(response.RawQuery)
	fmt.Println(response.Port())

	queryParams := response.Query()
	fmt.Printf("Type of q params: %T\n", queryParams)
	//fmt.Println(queryParams["coursename"])

	for _, value := range queryParams {
		fmt.Println(value)
	}

	queryParamsToBeUpdated := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=ron",
	}
	url2 := queryParamsToBeUpdated.String()
	fmt.Println(url2)
}