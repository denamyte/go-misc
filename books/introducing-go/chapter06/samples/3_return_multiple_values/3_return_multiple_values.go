package main

import "fmt"

func main() {
	fmt.Println(f())
	fmt.Println(f1())
}

func f() (int, int) {
	return 5, 6
}

func f1() (r1 int, r2 int) {
	r1 = 5
	r2 = r1 / 2
	return
}
