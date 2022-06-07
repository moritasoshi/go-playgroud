package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const (
		flag1 = 1 << iota
		flag2
		flag3
		flag4
		flag5
		flag6
		flag7
		flag8
		flag9
		flag10
	)

	fmt.Println(flag1, flag2, flag3, flag4, flag5, flag6, flag7, flag8, flag9, flag10)

	// if
	if a := flag1 == 1; a {
		fmt.Println(a, flag1)
	}

	// switch
	switch flag1 {
	case 1:
		fmt.Println("case 1")
	default:
		fmt.Println("default")
	}

	// generate random number
	t := time.Now().UnixNano()
	rand.Seed(t)
	for i := 0; i <= 20; i++ {
		n := rand.Intn(6)
		fmt.Print(n)
	}
	fmt.Println()

	// composite
	var ns []int
	var m map[string]int
	fmt.Println(ns)
	fmt.Println(m)

	// slice
	bufa := []string{}
	bufa = make([]string, 0, 15)
	for i := 0; i < 12; i++ {
		bufa = append(bufa, "Red")
		fmt.Println(len(bufa), cap(bufa))
	}

	// array
	bufb := [10]string{}
	for i := 0; i < 12; i++ {
		bufb[i] = "Red"
		fmt.Println(len(bufa), cap(bufa))
	}

	// var p struct {
	// 	name string
	// 	age  int
	// }

}
