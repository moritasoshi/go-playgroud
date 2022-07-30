package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: localhost\r\n\r\n")
	io.Copy(os.Stdout, conn)

	req, _ := http.NewRequest("GET", "http://localhost:8888", nil)
	a := req.Write(conn)
	fmt.Println("+++++++++")
	fmt.Print(a)

}
