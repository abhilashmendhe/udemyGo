package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, "is", p.age, "years old.")
}
func main() {
	p1 := person{
		first: "Abhi",
		age:   29,
	}
	p2 := person{
		first: "John",
		age:   42,
	}
	p1.speak()
	p2.speak()
}
