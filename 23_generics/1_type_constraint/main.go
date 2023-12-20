package main

import "fmt"

func addT[T int | float64](a, b T) T {
	return a + b
}
func main() {
	fmt.Println(addT(3, 2))
	fmt.Println(addT(3.2, 2.4))

}
