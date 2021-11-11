package main

import (
	"fmt"
	"time"
)

func main()  {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	createdDate := time.Date(2020, time.February, 23, 11, 12, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}