package main

import "fmt"

func main() {
	a := new(Android)
	a.Person.Talk()
	a.Talk()
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Model struct {
	SerialNumber string
}

type Android struct {
	Person   // an embedded type
	model Model
}




