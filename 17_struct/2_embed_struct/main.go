package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person Person
	ltk    bool
	first  string
}

func main() {

	sa1 := secretAgent{
		person: Person{
			first: "James",
			last:  "Bond",
			age:   42,
		},
		first: "etahn",
		ltk:   true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.ltk, sa1.person.first, sa1.first)
}
