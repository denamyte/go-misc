package main

import "fmt"

func main() {
	x := [15]float64{1, 2, 3, 4, 5, 6, 7.2, 8, 8.5, 9, 10}
	fmt.Println(x)
	s1 := x[:]
	fmt.Println(s1)
	s1 = x[:2]
	fmt.Println(s1)

}
