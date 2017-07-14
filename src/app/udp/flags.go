package main

import cli "gopkg.in/urfave/cli.v1"

var (
	LocalFlag = cli.StringFlag{
		Name:  "local",
		Usage: "local udp url",
		Value: ":40404",
	}
	UurlFlag = cli.StringFlag{
		Name:  "uurl",
		Usage: "remote udp url",
		Value: ":40404",
	}
	DataFlag = cli.StringFlag{
		Name:  "data",
		Usage: "the data to send",
	}
)

var udpCliCmd = cli.Command{
	Action:    udpCli,
	Name:      "udpcli",
	Usage:     "start a udp cli and send",
	ArgsUsage: "<udp remote url>",
	Category:  "UDP CLIENT",
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
	Description: `the udp cli`,
}

var udpSvrCmd = cli.Command{
	Action:    udpSvr,
	Name:      "udpsvr",
	Usage:     "start a udp server",
	ArgsUsage: "<udp local url>",
	Category:  "UDP SERVER",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "local",
			Usage: "local udp url",
			Value: ":40404",
		},
	},
	Description: `the udp server`,
}
