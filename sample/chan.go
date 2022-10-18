package main

import "fmt"

func main() {
	type s struct {
		val string
		err error
	}
	ch := make(chan string)

	// go func() {
	// 	ch <- "hello"
	// 	close(ch)
	// 	// ch <- "world"
	// }()

	val := <-ch
	fmt.Printf("%v\n", val)

	val = <-ch
	fmt.Printf("%v\n", val)

}
