package main

import (
	"fmt"
)

func main() {
	x := powinator(3)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
func powinator(a int) func() int {
	var c int = 1
	return func() int {
		c *= a
		return c
	}
}

// func powinator(a float64) func() float64 {
// 	var c float64 = 0
// 	return func() float64 {
// 		c++
// 		return math.Pow(a, c)
// 	}
// }
