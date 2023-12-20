package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(val int) int {
	if val == 1 {
		return 1
	}
	return val * factorial(val-1)
}
