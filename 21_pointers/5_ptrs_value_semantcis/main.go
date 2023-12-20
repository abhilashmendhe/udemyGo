package main

import "fmt"

// this is value semantics
func addOne(v int) int {
	return v + 1
}

// pointers semantic
func addOneP(v *int) {
	*v += 1
}
func main() {
	// value semantics
	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))

	// ptrs semantics
	b := 1
	fmt.Println("\nptrs semanics")
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
}
