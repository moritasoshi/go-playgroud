package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	fmt.Printf("data: %d\n", data)
	var i int32

	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)

	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 10000)
	fmt.Printf("bs: %d\n", bs)
	var i2 int32

	binary.Read(bytes.NewReader(bs), binary.BigEndian, &i2)
	fmt.Printf("bs: %d\n", i2)
}
