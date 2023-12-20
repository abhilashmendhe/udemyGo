package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	sum := foo(arr...)
	fmt.Println(sum)
	fmt.Println(bar(arr))
}
func foo(arr ...int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func bar(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
