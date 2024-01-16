package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Print("Enter a message: ")
		reader := bufio.NewReader(os.Stdin)
		sendstr, _ := reader.ReadString('\n')
		io.WriteString(conn, "echo>> "+sendstr)
		conn.Close()
	}
}
