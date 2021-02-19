package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("books/introducing-go/chapter08/samples/4_file_readdir")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = dir.Close() }()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
