package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("ok")
	ln, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConnection(conn)
	}

}
func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		log.Println("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}
