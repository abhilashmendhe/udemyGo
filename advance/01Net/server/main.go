package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("Listening on port 8000")
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	// for {
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, fmt.Sprint("Hello World\n", time.Now(), "\n"))
	conn.Close()
	// }
}
