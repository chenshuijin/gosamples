package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	cli "gopkg.in/urfave/cli.v1"
)

var app = cli.NewApp()

func init() {

	app.Name = filepath.Base(os.Args[0])
	app.Author = "csj"
	app.Email = "785795635@qq.com"
	app.Version = "1.0"
	app.Usage = "the udp command line interface"
	app.Commands = []cli.Command{
		udpCliCmd,
		udpSvrCmd,
	}
}

func main() {
	fmt.Println("begin")
	log.Println("gogogogo...")
	if err := app.Run(os.Args); err != nil {
		log.Fatal("app run err:%v", err)
	}

}
