package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		duration := time.Duration(rand.Intn(250))
		time.Sleep(duration * time.Millisecond)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	_, _ = fmt.Scanln(&input)
}
