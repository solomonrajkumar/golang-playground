package main

import "fmt"

func main()  {
	ronith := Person{"Ronith", 4, "ronith.r@yopmail.com"}
	fmt.Println(ronith)
	fmt.Printf("Individual values: %+v\n", ronith)
	fmt.Printf("Name: %v, Age: %v\n", ronith.Name, ronith.Age)
	ronith.GetName()
	ronith.NewEmail()
	fmt.Printf("Individual values: %+v\n", ronith)

}

type Person struct {
	Name  string
	Age   int
	Email string
}

func (person Person) GetName() {
	fmt.Println("Name:", person.Name)
}

func (person *Person) NewEmail() {
	person.Email = "ron@mailinator.com"
}