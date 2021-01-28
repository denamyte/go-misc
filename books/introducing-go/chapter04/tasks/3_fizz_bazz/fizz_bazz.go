package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%s ", getStr(i))
	}
}

func getStr(i int) string {
	if i%3 == 0 {
		if i%5 == 0 {
			return "FizzBazz"
		}
		return "Fizz"
	}
	if i%5 == 0 {
		return "Bazz"
	}
	return strconv.Itoa(i)
}