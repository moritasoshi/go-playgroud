package main

import (
	"fmt"
	"time"
)

func main() {
	sec := 10

	fmt.Printf("started at  %v\n", time.Now())
	t := <-time.After(time.Duration(sec) * time.Second)
	fmt.Printf("finished at %v\n", t)

}
