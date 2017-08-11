package main

import cli "gopkg.in/urfave/cli.v1"

var (
	LocalFlag = cli.StringFlag{
		Name:  "local",
		Usage: "local udp url",
		Value: ":40404",
	}
	UurlFlag = cli.StringFlag{
		Name:  "url",
		Usage: "remote udp url",
		Value: ":40404",
	}
	DataFlag = cli.StringFlag{
		Name:  "data",
		Usage: "the data to send",
		Value: "1",
	}
)
