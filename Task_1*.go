package main

import (
	"fmt"
)

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Привет, дорогой друг: %s!", p.Name)
}

func GreetSomeone(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	person := Person{Name: "Марк"}
	GreetSomeone(person)
}
