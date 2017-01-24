package main

import (
	"net"
)

type Queue struct {
}

type Item struct {
	Local  net.Conn
	Remote net.Conn
	Data   []byte
}
