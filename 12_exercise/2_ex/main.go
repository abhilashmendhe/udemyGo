package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x: %v.\t", x)

	if x >= 0 && x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Less than 250")
	}
}
