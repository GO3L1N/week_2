package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) Speak() {
	fmt.Println("Издает звук: ")
}

type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

func main() {
	animal := Animal{name: "Random animal"}
	animal.Speak()
	dog := Dog{Animal{name: "Buddy"}}
	dog.Speak()
}
