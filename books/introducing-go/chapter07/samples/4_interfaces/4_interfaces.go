package main

import (
	"fmt"
	"math"
)

func main() {
	shapes := []Shape{
		&Circle{0, 0, 2},
		&Rectangle{0, 0, 10.0, 5.0},
	}
	fmt.Println(totalArea(shapes...))

	multiShape := MultiShape{
		shapes: shapes,
	}
	fmt.Println(totalArea(&multiShape))
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x, y, width, height float64
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) perimeter() float64 {
	var per float64
	for _, v := range m.shapes {
		per += v.perimeter()
	}
	return per
}

func (m *MultiShape) area() float64 {
	return totalArea(m.shapes...)
}

func totalArea(shapes ...Shape) float64 {
	var total float64
	for _, v := range shapes {
		total += v.area()
	}
	return total
}

