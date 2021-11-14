package main

import (
	"fmt"
	"sync"
)

func main()  {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)
	//wg.Add(1)
	go func(w *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("Routine 1")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		w.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(w *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("Routine 2")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		w.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(w *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("Routine 3")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		w.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}