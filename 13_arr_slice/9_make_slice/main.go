package main

import (
	"fmt"
)

func main() {
	si := []string{"a", "b", "c"}
	fmt.Println(si)
	si[1] = "d"
	fmt.Println(si)

	xi := make([]int, 0, 2)
	fmt.Println(xi)
	fmt.Println("Capacity: ", cap(xi))
	fmt.Println("Length: ", len(xi))
	fmt.Println("-------------------")
	xi = append(xi, 100)
	xi = append(xi, 111)
	fmt.Println(xi)
	fmt.Println("Capacity: ", cap(xi))
	fmt.Println("Length: ", len(xi))
	fmt.Println("-------------------")
	xi = append(xi, 200)
	// xi = append(xi, 211)
	fmt.Println(xi)
	fmt.Println("Capacity: ", cap(xi))
	fmt.Println("Length: ", len(xi))
	fmt.Println("-------------------")
	xi = append(xi, 1001)
	xi = append(xi, 1111)
	fmt.Println(xi)
	fmt.Println("Capacity: ", cap(xi))
	fmt.Println("Length: ", len(xi))
}
