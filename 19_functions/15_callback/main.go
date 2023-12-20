package main

import "fmt"

func main() {
	add := doMath(10, 20, add)
	fmt.Println(add)
	sub := doMath(20, 14, sub)
	fmt.Println(sub)
}
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
