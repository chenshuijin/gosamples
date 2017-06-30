package main

import (
	"log"
	"net"

	cli "gopkg.in/urfave/cli.v1"
)

var udpCliCmd = cli.Command{
	Action:    udpCli,
	Name:      "udpcli",
	Usage:     "start a udp cli and send",
	ArgsUsage: "<udp remote url>",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "uurl",
			Usage: "remote udp url",
			Value: ":40404",
		},
		cli.StringFlag{
			Name:  "data",
			Usage: "the data to send",
		},
	},
	Category:    "udp cli command",
	Description: `the udp cli`,
}

func udpCli(ctx *cli.Context) error {
	remoteUrl := ctx.Args().First()
	log.Println("remote url:", remoteUrl)

	addr, err := net.ResolveUDPAddr("udp", ":40404")
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
