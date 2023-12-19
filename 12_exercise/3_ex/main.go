package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x: %v.\t", x)
	switch {
	case x <= 100:
		fmt.Println("Less than 100")
	case x <= 200:
		fmt.Println("Less than 200")
	default:
		fmt.Println("Less than 250")
	}
}
