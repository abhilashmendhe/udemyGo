package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%v\t%T\n", &x, &x)
	y := &x
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)

	*y = 43
	fmt.Println(x)

}
