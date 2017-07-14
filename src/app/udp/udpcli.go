package main

import (
	"log"
	"net"

	cli "gopkg.in/urfave/cli.v1"
)

func udpCli(ctx *cli.Context) error {
	log.Println("remote url:", ctx.GlobalString("uurl"))

	addr, err := net.ResolveUDPAddr("udp", ctx.GlobalString("uurl"))
	if err != nil {
		log.Fatal("Resolve udp addr err:", err)
	}
	cn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Dail udp err:", err)
	}
	defer cn.Close()
	b := []byte{0x01}
	count, err := cn.Write(b)
	if err != nil {
		log.Fatal("write to udp err:", err)
	}
	log.Println("write amount :", count)
	return nil
}
