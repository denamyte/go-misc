package main

import (
	"fmt"
	"math"
)

func main() {
	var c Circle  // an instance of a zero Circle
	fmt.Println(c)

	c2 := new(Circle)  // a pointer to a zero Circle
	fmt.Println(c2)

	c3 := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c3)

	c4 := Circle{0, 0, 5}
	fmt.Println(c4)

	pCircle := &Circle{0, 0, 5}
	fmt.Println(pCircle)

	fmt.Println(circleArea(c4))

	fmt.Println(circleAreaPtr(&c4))
	fmt.Println(circleAreaPtr(pCircle))
}

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaPtr(c *Circle) float64 {
	return math.Pi * c.r * c.r

}
