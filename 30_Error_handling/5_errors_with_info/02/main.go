package main

import (
	"fmt"
	"log"
	"math"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}
func main() {

	v, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(v)
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: Square root of negative number")
		return 0, norgateMathError{"50.234 N", "99.3234 W", nme}
	}
	return math.Sqrt(f), nil
}
