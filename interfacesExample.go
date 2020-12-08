package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type rectangular struct {
	width, height float64
}

func (r rectangular) area() float64 {
	return r.width * r.height
}

func (r rectangular) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Printf("The type is: %T\n", g)
	fmt.Println("The area is:", g.area())
	fmt.Println("The perimeter is:", g.perimeter())
}

func main() {
	// circle1 := circle{radius: 3}
	// rectangular1 := rectangular{width: 3, height: 4}

	// measure(circle1)
	// measure(rectangular1)

	var c geometry = circle{radius: 4}
	var r geometry = rectangular{width: 2, height: 3}
	measure(c)
	measure(r)
}
