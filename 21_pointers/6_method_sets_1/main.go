package main

import "fmt"

type Dog struct {
	first string
}

func (d Dog) Walk() {
	fmt.Println("My name is ", d.first, "and I'm walking")
}
func (d *Dog) Run() {
	fmt.Println("My name is ", d.first, "and I'm running")
}
func main() {
	d1 := Dog{
		first: "Henry",
	}
	d1.Walk()
	d1.Run()
	d2 := &Dog{
		first: "Padget",
	}
	d2.Walk()
	d2.Run()
}
