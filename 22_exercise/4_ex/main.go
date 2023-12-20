package main

import "fmt"

type Person struct {
	first string
}

func main() {
	p := Person{
		first: "Henry",
	}
	fmt.Println(p)
	p = changeName(p, "Hardy")
	fmt.Println(p)

	p2 := Person{
		first: "abhi",
	}
	fmt.Println(p2)
	changeNameP(&p2, "Abhilash")
	fmt.Println(p2)
}

// value semantics
func changeName(p Person, s string) Person {
	p.first = s
	return p
}

// ptrs semantics
func changeNameP(p *Person, s string) {
	p.first = s
}
