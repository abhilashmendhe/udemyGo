package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.first)
}
func saySomething(h human) {
	h.speak()
}
func main() {
	p := person{
		first: "abhi",
	}
	// p.speak()
	saySomething(&p)
}
