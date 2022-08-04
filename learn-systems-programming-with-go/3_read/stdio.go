package main

import (
	"fmt"
	"io"
	"os"
)

// os.Stdin fulfills
// the io.Reader and io.Closer interface
func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s' \n", size, string(buffer))
	}
}
