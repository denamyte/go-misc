package main

import "fmt"

func greatest(params ...int) (g int) {
	if len(params) > 0 {
		g = params[0]
	}
	for _, v := range params {
		if v > g {
			g = v
		}
	}
	return
}

func main() {
	fmt.Println(greatest(1, 2, 3, 4, 5, 2, -1))
}
