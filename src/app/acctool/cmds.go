package main

import cli "gopkg.in/urfave/cli.v1"

var (
	loopKeyscmd = cli.Command{
		Action:      LoopKeys,
		Name:        "LoopKeys",
		Category:    "Generate Keys",
		Description: "Loop Generate Keys",
	}
)
