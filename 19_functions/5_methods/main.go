package main

import "fmt"

type Person struct {
	first string
}

func (p Person) speak() {
	fmt.Println("I am", p.first)
}
func main() {
	p1 := Person{
		first: "abhi",
	}
	p2 := Person{
		first: "jenny",
	}
	p1.speak()
	p2.speak()
}
