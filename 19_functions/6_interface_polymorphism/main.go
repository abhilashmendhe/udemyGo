package main

import "fmt"

type Person struct {
	first string
}
type secretAgent struct {
	Person
	ltk bool
}

func (p Person) speak() {
	fmt.Println("I am", p.first)
}
func (sa secretAgent) speak() {
	fmt.Println("I am secret agent:", sa.first)
}

// Since both type secretAgent and Person have methods speak then they are also
// of type human. human interface has speak method. This is Polymorphism
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	sa := secretAgent{
		Person: Person{
			first: "abhi",
		},
		ltk: true,
	}
	p2 := Person{
		first: "jenny",
	}
	// sa.speak()
	// p2.speak()
	saySomething(sa)
	saySomething(p2)
}
