package main

import "fmt"

func main() {

	sl := make([]int, 10)
	ind := 0
	for i := 42; i < 52; i++ {
		sl[ind] = i
		ind++
	}
	fmt.Println(sl)
}
