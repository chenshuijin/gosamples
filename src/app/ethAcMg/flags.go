package main

import cli "gopkg.in/urfave/cli.v1"

var (
	PassFlag = cli.StringFlag{
		Name:  "p",
		Usage: "password of account",
		Value: "",
	}
	KeystoreFlag = cli.StringFlag{
		Name:  "k",
		Usage: "keystore of eth accounts",
		Value: "~/.ethkeystore",
	}
)
