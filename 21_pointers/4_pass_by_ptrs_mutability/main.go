package main

import "fmt"

func intDelta(n *int) {
	*n = 1023
}
func sliceDelta(li []int) {
	li[0] = 99
}
func mapDelta(m map[string]int, s string) {
	m[s] = 32
}
func main() {
	a := 42
	fmt.Println("Value of a:", a)
	intDelta(&a)
	fmt.Println("Value of a:", a)

	// slice are pass by reference
	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	// map are also pass by reference
	m := make(map[string]int)
	m["james"] = 40
	m["abhi"] = 29
	fmt.Println(m)
	mapDelta(m, "james")
	fmt.Println(m)

}
