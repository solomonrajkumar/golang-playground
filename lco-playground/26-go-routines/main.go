package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually pointers

func main()  {
	//go greeter("raj")
	//greeter("ron")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}
//
//func greeter(s string) {
//	for i := 0; i < 5; i++ {
//		fmt.Println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	statusCode := result.StatusCode
	fmt.Printf("%d status code for %s\n", statusCode, endpoint)
}