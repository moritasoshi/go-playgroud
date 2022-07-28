package main

import "fmt"

type Talker interface {
	Talk()
}

type (
	Greeter struct {
		name string
	}
	Greeter1 struct {
		age int
	}
)

func (g Greeter) Talk() {
	fmt.Printf("Hello. my name is %s.\n", g.name)
}

func (g Greeter1) Talk() {
	fmt.Printf("Hello. I am %d years old.\n", g.age)
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()

	talker = &Greeter1{20}
	talker.Talk()
}
