package main

import "fmt"

func zero(x int) {
	x = 0
}

func zero2(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	zero2(&x)
	fmt.Println(x)
}
