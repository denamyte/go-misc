package main

import "fmt"

func main() {
	fmt.Print("Enters feet number: ")
	var feet float64
	_, _ = fmt.Scan(&feet)

	fmt.Printf("%.2f feet are %.2f meters", feet, feet * .3048)
}
