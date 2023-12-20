package main

import "fmt"

func main() {
	defer foo4()
	defer foo3()
	defer foo2()
	foo1()

}

func foo1() {
	fmt.Println("Hello 1")
}
func foo2() {
	fmt.Println("Hello 2")
}
func foo3() {
	fmt.Println("Hello 3")
}
func foo4() {
	fmt.Println("Hello 4")
}
