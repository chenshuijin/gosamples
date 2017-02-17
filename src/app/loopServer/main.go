package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var localHost = flag.String("local", ":8080", "the local listen host")

func main() {
	fmt.Println("ok")
	flag.Parse()
	startSvr()
	log.Println("stop")
}

func startSvr() {
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
			continue
		}
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("begin handle")
	for {
		data := make([]byte, 2048)
		count, err := conn.Read(data)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		log.Println("recv count:", count)
		if count == 3 && data[0] == 0xff && data[1] == 0xfb && data[2] == 0x06 {
			log.Println("stop loop!")
			break
		}
		count, err = conn.Write(data)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		log.Printf("write to local data:\n%s\n", data)
	}
	return
}
