package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Greet() {
	fmt.Println("Привет, дорогой друг:", p.name)
}

func main() {
	Maxim := Person{
		name:    "Максим",
		age:     20,
		address: "Москва",
	}
	Vasya := Person{
		name:    "Василий",
		age:     22,
		address: "Дубай",
	}
	Maxim.Greet()
	Vasya.Greet()
}
