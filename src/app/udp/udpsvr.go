package main

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	cli "gopkg.in/urfave/cli.v1"
)

func udpSvr(ctx *cli.Context) error {
	log.Println("udp svr...")
	log.Println("loacl:", ctx.String(LocalFlag.Name))
	nodekey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("could not generate key:%v", err)
		return err
	}
	if err = crypto.SaveECDSA("./nodekey", nodekey); err != nil {
		log.Fatal("%v", err)
		return err
	}
	u, err := ListenUDP(nodekey, ctx.String(LocalFlag.Name))
	go u.readLoop()
	for {
		select {
		case buf := <-MsgChan:
			log.Println("recv:", buf)
		}
	}
	return nil
}
