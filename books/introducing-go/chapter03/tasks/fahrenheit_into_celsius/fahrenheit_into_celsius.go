package main

import "fmt"

func main() {
	fmt.Print("Enter a number corresponding to Fahrenheit degrees: ")
	var fahrenheit float64
	_, _ = fmt.Scan(&fahrenheit)
	var celsius = 5.0 / 9 * (fahrenheit - 32)
	fmt.Println(celsius)
}
