package main

import "fmt"

//goland:noinspection GoBoolExpressions
func main() {
	fmt.Println((true && false) || (false && true) || ! (false && false))
}
