package main

import cli "gopkg.in/urfave/cli.v1"

var (
	loopKeyscmd = cli.Command{
		Action:      LoopKeys,
		Name:        "LoopKeys",
		Category:    "Generate Keys",
		Description: "Loop Generate Keys",
	}
	qbalcmd = cli.Command{
		Action:      qbal,
		Name:        "qbal",
		Category:    "Query balance of address",
		Description: "Query balance of address by call web api",
	}
)
