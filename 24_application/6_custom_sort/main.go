package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func main() {
	p1 := person{
		first: "abhi",
		age:   29,
	}
	p2 := person{
		first: "vicky",
		age:   25,
	}
	p3 := person{
		first: "himanshu",
		age:   31,
	}
	people := []person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
