package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("books/introducing-go/chapter08/samples/1_file/temp.txt")
	if err!= nil {
		// handle the error here
		fmt.Println(err)
		return
	}
	defer func() { _ = file.Close() }()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
