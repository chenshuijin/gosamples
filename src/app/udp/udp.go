package main

import (
	"crypto/ecdsa"
	"log"
	"net"
)

//var MsgQueue = tunnel.NewSafeQueue()
var MsgChan = make(chan []byte, 100)

type udp struct {
	conn *net.UDPConn
	priv *ecdsa.PrivateKey
}

func ListenUDP(priv *ecdsa.PrivateKey, laddr string) (*udp, error) {
	addr, err := net.ResolveUDPAddr("udp", laddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	log.Println("listen udp on:", addr.String())
	return &udp{conn: conn, priv: priv}, nil
}

func (u *udp) readLoop() {
	defer u.conn.Close()
	buf := make([]byte, 1280)
	for {
		nbytes, from, err := u.conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		} else {
			log.Printf("recv %d bytes from [%s].\n", nbytes, from)
			MsgChan <- buf
		}
	}
}
