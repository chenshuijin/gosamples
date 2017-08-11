package main

import (
	"log"
	"net"
	"sync"

	cli "gopkg.in/urfave/cli.v1"
)

func udpCli(ctx *cli.Context) error {
	log.Println("remote url:", ctx.String(UurlFlag.Name))
	addr, err := net.ResolveUDPAddr("udp", ctx.String(UurlFlag.Name))
	if err != nil {
		log.Fatal("Resolve udp addr err:", err)
	}
	cn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Dail udp err:", err)
	}
	log.Println("udp cli local:", cn.LocalAddr())
	defer cn.Close()
	data := ctx.String(DataFlag.Name)
	b := []byte(data)
	count, err := cn.Write(b)
	if err != nil {
		log.Fatal("write to udp err:", err)
	}
	log.Println("write amount :", count)
	wg := sync.WaitGroup{}
	wg.Add(1)
	//	go readLoop(cn)
	wg.Wait()
	return nil
}
