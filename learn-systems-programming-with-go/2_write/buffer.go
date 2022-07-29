package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var bb bytes.Buffer
	str := "bytes.Buffer Write example\n"

	bb.Write([]byte(str))
	fmt.Print(1, bb.String())

	bb.WriteString(str)
	fmt.Print(2, bb.String())

	io.WriteString(&bb, str)
	fmt.Print(3, bb.String())

	// builder
	fmt.Println()
	builder()
}

func builder() {
	var builder strings.Builder
	str := "bytes.Buffer Write example\n"

	builder.Write([]byte(str))
	fmt.Print(4, builder.String())
}
