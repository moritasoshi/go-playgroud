package main

import (
	"fmt"
	"reflect"
)

// このinterface内の全methodを実装していれば
// その構造体はTalker型で宣言できる
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

func (g Greeter1) Meow() {
	fmt.Println("Meow")
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()

	talker = &Greeter1{20}
	talker.Talk()
	// talker.Meow() // interfaceに無いメソッドは呼び出せない

	what := &Greeter1{100}
	what.Talk()
	what.Meow() // こっちは呼び出せる

	fmt.Printf("talker: %T\n", talker)
	fmt.Printf("what: %T\n", what)

	fmt.Println("Is same type?: ", reflect.TypeOf(talker) == reflect.TypeOf(what)) // true
}
