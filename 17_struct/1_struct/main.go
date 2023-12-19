package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := Person{
		first: "Abhilash",
		last:  "Mendhe",
		age:   29,
	}
	p2 := Person{
		first: "Sarah",
		last:  "Illustrates",
		age:   30,
	}

	// fmt.Println(p1)
	// fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T \t %#v\n", p2, p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
