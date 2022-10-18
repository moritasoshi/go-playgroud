package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func parallel(wg *sync.WaitGroup) {
	fmt.Println("博")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("多")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("の")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("塩")
	wg.Done()
}

func TestWaitGroup(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go parallel(wg)
	}
	wg.Wait()
}

func TestBroadcast(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	l := new(sync.Mutex)
	c := sync.NewCond(l)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("waiting %d\n", i)
			l.Lock()
			defer l.Unlock()
			c.Wait()
			fmt.Printf("go %d\n", i)
		}(i)
	}

	for i := 3; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	// 全てのgoroutineに一斉に号令をかける
	c.Broadcast()
	time.Sleep(3 * time.Second)
}

func TestBroadcastCond(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := new(sync.Cond)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("waiting %d\n", i)
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fmt.Printf("go %d\n", i)
		}(i)
	}

	for i := 3; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	// 全てのgoroutineに一斉に号令をかける
	c.Broadcast()
	time.Sleep(3 * time.Second)
}
