package main

import "fmt"

func main() {
	a := new(Android)
	fmt.Println(a.Talk())
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}
