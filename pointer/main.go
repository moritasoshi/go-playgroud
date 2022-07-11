package main

import "fmt"

func main() {
	fmt.Println("- 変数 T のポインタは、 *T 型で、ゼロ値は nil です。")
	fmt.Println("- & オペレータは、そのオペランドへのポインタを引き出します。")
	fmt.Println("- * オペレータは、ポインタの指す先の変数を示します。")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
