package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {

	v, err := sqrt(10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(v)
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: Square root of negative number")
	}
	return math.Sqrt(f), nil
}
