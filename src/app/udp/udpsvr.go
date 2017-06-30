package main

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	cli "gopkg.in/urfave/cli.v1"
)

var udpSvrCmd = cli.Command{
	Action:    udpSvr,
	Name:      "udpsvr",
	Usage:     "start a udp server",
	ArgsUsage: "<udp local url>",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "local",
			Usage: "local udp url",
			Value: ":40404",
		},
	},
	Category:    "udp server command",
	Description: `the udp server`,
}

func udpSvr(ctx *cli.Context) error {
	log.Println("udp svr...")
	nodekey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("could not generate key:%v", err)
		return err
	}
	if err = crypto.SaveECDSA("./nodekey", nodekey); err != nil {
		log.Fatal("%v", err)
		return err
	}
	u, err := ListenUDP(nodekey, ":40404")
	go u.readLoop()
	select {
	case buf := <-MsgChan:
		log.Println("recv:", buf)
	}
	return nil
}
