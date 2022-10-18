package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	// empty := make([]byte, 0)
	e := []byte{0}
	var b [8]byte

	copy(b[8-len(e):], e)
	data := binary.BigEndian.Uint64(b[:])
	fmt.Println(data)
}
