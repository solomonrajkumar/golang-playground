package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` // hides field
	Tags     []string `json:"tags,omitempty"`
}

func main()  {
	encodeJson()
}

func encodeJson() {
	lcocourses := []course{
		{"React", 23, "xyz.com", "abc123", []string{"Web", "JS"}},
		{"JS", 12, "abx.com", "qwe123", []string{"Dev", "TS"}},
		{"TS", 55, "wer.com", "zxc123", nil},
	}

	// package as JSON
	//finalJSON, err := json.Marshal(lcocourses)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%T\n", finalJSON)

	finalJSON, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}