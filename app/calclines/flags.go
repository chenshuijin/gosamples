package main

import cli "gopkg.in/urfave/cli.v1"

var (
	dirFlag = cli.StringFlag{
		Name:  "d",
		Usage: "The path of the code folder",
		Value: "./",
	}
	isAsyncFlag = cli.BoolFlag{
		Name:  "a",
		Usage: "Is to use multi-routine",
	}
)
