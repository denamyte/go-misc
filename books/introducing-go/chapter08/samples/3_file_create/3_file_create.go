package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("books/introducing-go/chapter08/samples/3_file_create/test_create.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = file.Close() }()

	_, _ = file.WriteString("test")
}
