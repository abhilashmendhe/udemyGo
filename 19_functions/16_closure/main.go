package main

import "fmt"

func main() {

	f := incrementor()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4

	g := incrementor() // now again x will start with 0 becuase g is at diff. mem location
	fmt.Println(g())   // 1
	fmt.Println(g())   // 2
	fmt.Println(g())   // 3
	fmt.Println(g())   // 4

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
