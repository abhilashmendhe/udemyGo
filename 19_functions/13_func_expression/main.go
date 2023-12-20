package main

import "fmt"

func main() {

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is anon. func. showing my name", s)
	}
	x()
	y("Abhilash")
}
