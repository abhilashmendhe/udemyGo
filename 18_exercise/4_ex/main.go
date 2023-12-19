package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Abhi",
		friends: map[string]int{
			"Jenny": 27,
			"Q":     70,
		},
		favDrinks: []string{"Pinacolada", "Rum"},
	}

	fmt.Println(p1)
}
