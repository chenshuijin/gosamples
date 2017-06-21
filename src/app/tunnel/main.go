package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
	"time"
)

var localHost = flag.String("local", ":8081", "the local listen host")
var remoteHost = flag.String("remote", "127.0.0.1:8080", "the remote host")

func main() {
	fmt.Println("ok")
	flag.Parse()
	switch runtime.GOOS {
	case "windows":
		log.Println("start a normal tunnel")
		startSvr()
	case "darwin":
		fallthrough
	default:
		log.Println("start a gracefull tunnel")
		startGraceSvr()
	}

	log.Println("stop")
}

func startGraceSvr() {
	gs, err := New("tcp", *localHost, *remoteHost)
	if err != nil {
		log.Fatal("new listener err:", err)
	}
	gs.Start()
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
	squeue = NewSafeQueue()
	sqChan = make(chan interface{}, 100)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

type Tunnel struct {
	From *net.Conn
	To   *net.Conn
	Data []byte
}

var TimeOut time.Time
var squeue *SafeQueue
var sqChan chan interface{}

func handleConnection(conn net.Conn) {
	remoteConn, err := net.Dial("tcp", *remoteHost)
	if err != nil {
		fmt.Println("err:", err)
	}
	go Read2Queue(conn, remoteConn)
	go Read2Queue(remoteConn, conn)
	go WriteFromQueue()

	return
}

func Read2Queue(from, to net.Conn) {
	defer from.Close()
	defer to.Close()
	from.SetDeadline(TimeOut)
	to.SetDeadline(TimeOut)

	for {
		data := make([]byte, 2048)
		count, err := from.Read(data)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		log.Println("recv count:", count)
		squeue.Push(&Tunnel{
			From: &from,
			To:   &to,
			Data: data[:count],
		})
		sqChan <- 1
		log.Println("read2queue sqchan")
	}
}

func WriteFromQueue() {
	for {
		select {
		case t := <-sqChan:
			log.Println("<- from sqchain:", t)
			tmp := squeue.Pop().(*Tunnel)
			count, err := (*tmp.To).Write(tmp.Data)
			if err != nil {
				log.Println("err:", err)
				(*tmp.To).Close()
				(*tmp.From).Close()
			}
			log.Println("write count:", count)
		}
	}
}
