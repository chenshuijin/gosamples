package main

import cli "gopkg.in/urfave/cli.v1"

var (
	ConfFlag = cli.StringFlag{
		Name:  "c",
		Usage: "Config file path",
		Value: "./config.json",
	}
)
