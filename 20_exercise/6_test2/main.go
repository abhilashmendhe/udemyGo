package main

import "fmt"

func main() {
	fmt.Printf("%T\n", doMath)
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)

}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}
