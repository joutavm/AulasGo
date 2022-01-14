package formas

import (
	"math"
)

type Geometry interface {
	area() float64
}

func GetArea(g Geometry) float64 {
	return g.area()
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
