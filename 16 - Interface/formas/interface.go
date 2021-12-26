package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

func getArea(g geometry) float64 {
	return g.area()
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {

	fmt.Println(getArea(rectangle{width: 10, height: 20}))
	fmt.Println(getArea(circle{10}))

}
