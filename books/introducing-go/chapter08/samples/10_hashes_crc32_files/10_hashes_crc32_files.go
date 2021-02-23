package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	filename1 := "books/introducing-go/chapter08/samples/10_hashes_crc32_files/test_file_1.txt"
	filename2 := "books/introducing-go/chapter08/samples/10_hashes_crc32_files/test_file_2.txt"
	h1, err := getHash(filename1)
	if err != nil {
		return
	}
	h2, err := getHash(filename2)
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}

func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer f.Close()
	h := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(h, f)
	// we don't care about how many bytes were written, but we do want to
	// handle the error
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}
