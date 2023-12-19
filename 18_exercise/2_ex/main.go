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
	p2 := Person{
		first: "Abhi",
		last:  "Menda",
		favIC: []string{"Strawberry", "Pineapple"},
	}

	m1 := make(map[string]Person)

	m1["Bond"] = p1
	m1["Mendhe"] = p2

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
