package main

import (
	"fmt"
)

type Mover interface {
	Move()
}
type Flyer interface {
	Fly()
}
type Superhero struct {
	Name string
}

func (s Superhero) Move() {
	fmt.Printf("%s is moving!\n", s.Name)
}
func (s Superhero) Fly() {
	fmt.Printf("%s is flying!\n", s.Name)
}
func performAction(m Mover, f Flyer) {
	m.Move()
	f.Fly()
}

func main() {
	hero := Superhero{Name: "Superman"}

	performAction(hero, hero)
}
