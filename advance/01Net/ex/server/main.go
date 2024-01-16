package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	bs, _ := io.ReadAll(conn)
	fmt.Println(string(bs))
}
