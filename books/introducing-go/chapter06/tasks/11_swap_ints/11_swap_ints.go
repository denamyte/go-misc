package main

import "fmt"

func swapInts(iPtr1 *int, iPtr2 *int) {
	*iPtr1, *iPtr2 = *iPtr2, *iPtr1
}

func main() {
	i1, i2 := 1, 2
	swapInts(&i1, &i2)
	fmt.Println(i1, i2)
}
