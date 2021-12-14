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
	//decodeJSON()
}

func encodeJson() {
	lcocourses := []course{
		{"React", 23, "xyz.com", "abc123", []string{"Web", "JS"}},
		{"JS", 12, "abx.com", "qwe123", []string{"Dev", "TS"}},
		{"TS", 55, "wer.com", "zxc123", nil},
	}

	// package as JSON
	finalJSON, err := json.Marshal(lcocourses)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%T\n", finalJSON)
	fmt.Println(string(finalJSON))

	//finalJSON, err := json.MarshalIndent(lcocourses, "", "\t")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", finalJSON)
}

func decodeJSON() {
	sampleJson := []byte(`
		{
			"course_name": "React",
			"price": 23,
			"platform": "xyz.com",
			"tags": [
					"Web",
					"JS"
			]
        }
	`)
	var lcoCourse course
	checkValid := json.Valid(sampleJson)
	if checkValid {
		fmt.Println("VALID JSON")
		err := json.Unmarshal(sampleJson, &lcoCourse)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", lcoCourse)
		fmt.Println("@@@@@@@")
		fmt.Println(lcoCourse.Name)
		fmt.Printf("%T\n", lcoCourse.Name)
		fmt.Println(lcoCourse.Price)
		fmt.Printf("%T\n", lcoCourse.Price)
		fmt.Println(lcoCourse.Platform)
		fmt.Printf("%T\n", lcoCourse.Platform)
		fmt.Println(lcoCourse.Password)
		fmt.Printf("%T\n", lcoCourse.Password)
		fmt.Println(lcoCourse.Tags)
		fmt.Printf("%T\n", lcoCourse.Tags)
	} else {
		fmt.Println("INVALID JSON")
	}


	// parse and push to K-V pairs
	//var myData map[string]interface{}
	//json.Unmarshal(sampleJson, &myData)
	//fmt.Printf("%#s\n", myData)
	//
	//for k,v := range myData {
	//	fmt.Printf("%v : %v  :: %T\n", k, v, v)
	//}
}