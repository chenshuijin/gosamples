package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var localHost = flag.String("local", ":8080", "the local listen host")
var remoteHost = flag.String("remote", "127.0.0.1:7050", "the remote host")

func main() {
	fmt.Println("ok")
	flag.Parse()
	addr, err := net.ResolveTCPAddr("tcp", *localHost)
	if err != nil {
		panic(err)
	}

	ln, err := net.ListenTCP("tcp", addr)
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
	defer conn.Close()
	addr, err := net.ResolveTCPAddr("tcp", *remoteHost)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	remoteConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer remoteConn.Close()

	for {
		data := make([]byte, 2048)
		count, err := conn.Read(data)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		log.Println("recv count:", count)
		log.Printf("recv data:\n%s\n", data)
		count, err = remoteConn.Write(data)
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		log.Println("send count:", count)

		count, err = remoteConn.Read(data)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		log.Println("read from remote count:", count)
		log.Printf("read from remote data:\n%s\n", data)
		count, err = conn.Write(data)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		log.Println("write to local count:", count)
		log.Printf("write to local data:\n%s\n", data)
	}
	//	fmt.Fprintf(conn, "OK\n")
	return
}
