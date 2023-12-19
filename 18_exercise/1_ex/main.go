package main

import "fmt"

type Person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := Person{
		first: "John",
		last:  "Bond",
		favIC: []string{"Chocolate", "Butterscotch", "Vanilla"},
	}
	fmt.Println(p1)
}
