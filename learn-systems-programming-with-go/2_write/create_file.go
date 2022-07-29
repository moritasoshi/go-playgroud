package main

import (
	"log"
	"os"
)

func main() {

	if _, err := os.Stat("tmp/"); os.IsNotExist(err) {
		if err := os.Mkdir("tmp/", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create("tmp/create_file.txt")
	if err != nil {
		panic(err)
	}

	file.Write([]byte("os.File examples\n"))
	file.Close()
}
