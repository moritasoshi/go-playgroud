package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func heavy(i int) {
	fmt.Println(i)
	time.Sleep(5 * time.Second)
}

func TestMain(m *testing.M) {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(10000)
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			heavy(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(runtime.NumCPU())

}
