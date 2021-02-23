package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	_ = filepath.Walk("books/introducing-go/chapter08",
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})
}
