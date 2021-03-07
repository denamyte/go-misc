package main

import "fmt"
import m "github.com/denamyte/go-misc/books/introducing-go/chapter08/samples/16_creating_packages/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	min := m.Min(xs)
	max := m.Max(xs)
	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Min: %.2f\n", min)
	fmt.Printf("Max: %.2f\n", max)
}
