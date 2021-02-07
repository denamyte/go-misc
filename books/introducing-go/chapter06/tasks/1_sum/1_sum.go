package main

import "fmt"

func sum(ar []int) (s int) {
	for _, v := range ar {
		s += v
	}
	return
}

func main() {
	ar := []int{1, 2, 3, 4}
	fmt.Println(sum(ar))
}
