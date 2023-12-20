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

type youngin interface {
	Walk()
	Run()
}

func youngRun(y youngin) {
	y.Run()
}
func youngWalk(y youngin) {
	y.Walk()
}

func main() {
	d1 := Dog{
		first: "Henry",
	}
	d1.Walk()
	// youngRun(d1)
	// youngWalk(d1)

	d2 := &Dog{
		first: "Padget",
	}
	youngRun(d2)
	youngWalk(d2)
}
