package main

import "fmt"

func main() {

	xs := []int{43, 44, 45}
	fmt.Println(xs)

	xs = append(xs, 100)
	fmt.Println(xs)

	xs = append(xs, 99, 91, 23, 532)
	fmt.Println(xs)
}
