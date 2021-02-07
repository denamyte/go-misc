package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

func main() {
	oddGen := makeOddGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(oddGen())
	}
}
