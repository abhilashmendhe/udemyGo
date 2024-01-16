package main

import "fmt"

/* L - Liskov Substitution Principle (LSP)

 */

// Let's Consider a struct Animal
type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Println("Animal Sound")
}

/*
	 Now, let's say we want to create a new struct Bird that represents
		a specific type of animal
*/
type Bird struct {
	Animal
}

func (b Bird) MakeSound() {
	fmt.Println("Chirp chirp")
}

/* This principle states that objects of a superclass should be replaceable with
objects of a subclass without affecting the correctness of the program.
This helps to ensure that the relationships between classes are well-defined and maintainable.
*/

type AnimalBehavior interface {
	MakeSound()
}

// MakeSound represent a program that works with animals and is expected
// to work with base class (Animal) or any subclass (Bird in this case)
func MakeSound(ab AnimalBehavior) {
	ab.MakeSound()
}

func main() {
	a := Animal{}
	b := Bird{}
	MakeSound(a)
	MakeSound(b)
}

/*
This demonstrates inheritance in Go, as well as the Liskov Substitution Principle,
as objects of a subtype Bird can be used wherever objects of the base type Animal
are expected, without affecting the correctness of the program.
*/
