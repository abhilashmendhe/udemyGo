package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("This is anon. func. showing my name", s)
	}("Abhilash")
}
