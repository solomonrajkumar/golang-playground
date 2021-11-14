package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("Channels Demo")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}
	//myChannel <- 5
	//fmt.Println(<-myChannel)
	wg.Add(2)
	go func(ch chan int, group *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan int, group *sync.WaitGroup) {
		ch <- 2
		ch <- 4
		close(ch)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}