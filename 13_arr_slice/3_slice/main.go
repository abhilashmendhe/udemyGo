package main

import "fmt"

func main() {
	xs := []string{"hello", "world"}

	fmt.Println(xs)
	fmt.Printf("Type is %T\n", xs)

	/*
		slice is a DS with three elements
		-- len
		-- cap
		-- pointer to an underlying array

		type slice struct{
			array unsafe.pointer
			len int
			cap int
		}

	*/
}
