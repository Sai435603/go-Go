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

func Check(v interface{}) {
    switch x := v.(type) {

    case int:
        fmt.Println("Integer", x)

    case string:
        fmt.Println("String", x)

    case bool:
        fmt.Println("Boolean", x)

    default:
        fmt.Println("Unknown")
    }
}

func main() {
	rect := Rectangle{Length: 5, Width: 3}
	circle := Circle{Radius: 2}

	fmt.Printf("Rectangle area: %.2f\n", rect.Area())
	fmt.Printf("Circle area: %.2f\n", circle.Area())
	Check(69)
}
