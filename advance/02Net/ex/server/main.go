package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var KV = make(map[string]string)

func handle(conn net.Conn, ln net.Listener) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		txt := scanner.Text()
		s_arr := strings.Split(txt, " ")
		cond := s_arr[0]
		key := s_arr[1]

		switch cond {
		case "GET":
			val, is := KV[key]
			if is {
				fmt.Println(key, "->", val)
			} else {
				fmt.Println(key, "not found!")
			}
		case "SET":
			if len(s_arr) > 2 {
				val := s_arr[2]
				KV[key] = val
				fmt.Println(key, "->", val, "is set")
			} else {
				fmt.Println("Please specify value!")
			}
		case "DEL":
			_, is := KV[key]
			if is {
				delete(KV, key)
				fmt.Println(key, "deleted!")
			} else {
				fmt.Println("Key not found to be deleted!")
			}
		case "CLOSE":
			conn.Close()
			ln.Close()
		default:
			fmt.Println("Wrong Input!!")
		}

	}
}
func main() {

	fmt.Println("Listening on 9000")
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn, ln)
	}
}
