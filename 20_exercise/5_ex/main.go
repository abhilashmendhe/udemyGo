package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.length * s.width
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Shape interface {
	area() float64
}

func info(s Shape) float64 {
	return s.area()
}

func main() {

	square := Square{
		length: 4.0,
		width:  4.0,
	}
	circle := Circle{
		radius: 10.0,
	}
	fmt.Println("Area of square:", info(square))
	fmt.Println("Area of circle:", info(circle))
}
