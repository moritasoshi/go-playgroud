package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	writeSize, err := io.Copy(newFile, oldFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("writeSize: %d\n", writeSize)

}
