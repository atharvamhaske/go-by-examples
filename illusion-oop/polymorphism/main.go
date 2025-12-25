package main

import (
    "fmt"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func printArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    circle := Circle{radius: 5}
    rectangle := Rectangle{width: 10, height: 5}

    printArea(circle)
    printArea(rectangle)
}

//Circle implements Shape interface because it implements its methods
//→ Circle.Area()
// an interface in Go is satisfied by any type that implements its methods