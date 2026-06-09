package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length, Width float64
}

type Circle struct {
	Radius float32
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width

}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func main() {
	rect := Rectangle{Length: 5, Width: 3}
	circle := Circle{Radius: 2}

	fmt.Printf("Rectangle area: %.5f\n", rect.Area())
	fmt.Printf("Circle area: %.5f\n", circle.Area())
}
