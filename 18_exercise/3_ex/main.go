package main

import "fmt"

type Engine struct {
	electric bool
}
type Vehicle struct {
	Engine
	make  string
	model string
	doors int
	color string
}

func main() {

	v1 := Vehicle{
		Engine: Engine{
			electric: true,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "Blue",
	}

	v2 := Vehicle{
		Engine: Engine{
			electric: false,
		},
		make:  "Ferrai",
		model: "F430",
		doors: 2,
		color: "Yellow",
	}

	fmt.Println(v1)
	fmt.Println(v2)
}
