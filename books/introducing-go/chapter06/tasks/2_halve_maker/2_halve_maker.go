package main

import "fmt"

func halveMaker(v int) (int, bool) {
	return v / 2, v % 2 == 0
}

func main() {
	params := []int {1, 2, 3, 4}
	for _, v := range params {
		r, b := halveMaker(v)
		fmt.Printf("param = %d, result: (%d, %t)\n", v, r, b)
	}
}
