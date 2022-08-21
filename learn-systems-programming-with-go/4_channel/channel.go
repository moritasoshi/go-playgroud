package main

import (
	"fmt"
	"math"
)

func primaNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primaNumber()

	// チャネルに値が入るたびにループが回る
	for n := range pn {
		fmt.Println(n)
	}
}
