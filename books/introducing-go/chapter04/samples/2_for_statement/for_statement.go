package main

import "fmt"

func main() {
	for1()
	for2()
}

func for1() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func for2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
