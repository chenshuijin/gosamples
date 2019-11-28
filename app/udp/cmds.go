package main

import cli "gopkg.in/urfave/cli.v1"

var (
	udpCliCmd = cli.Command{
		Action:    udpCli,
		Name:      "udpcli",
		Usage:     "start a udp cli and send",
		ArgsUsage: "<udp remote url>",
		Category:  "UDP CLIENT",
		Flags: []cli.Flag{
			UurlFlag,
			DataFlag,
		},
		Description: `the udp cli`,
	}

	udpSvrCmd = cli.Command{
		Action:    udpSvr,
		Name:      "udpsvr",
		Usage:     "start a udp server",
		ArgsUsage: "<udp local url>",
		Category:  "UDP SERVER",
		Flags: []cli.Flag{
			LocalFlag,
		},
		Description: `the udp server`,
	}
)
