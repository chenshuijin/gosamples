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
	PublicKeyFlag = cli.StringFlag{
		Name:  "pub",
		Usage: "public key bytes uncompressed form specified in section 4.3.6 of ANSI X9.62.",
		Value: "~/.ethkeystore",
	}
)
