package main

import cli "gopkg.in/urfave/cli.v1"

var (
	newEthAccountCmd = cli.Command{
		Action:    newEthAccount,
		Name:      "new",
		Usage:     "create new eth account under keystore",
		ArgsUsage: "",
		Category:  "NEW ETH ACCOUNT",
		Flags: []cli.Flag{
			PassFlag,
			KeystoreFlag,
		},
		Description: `create new eth account under keystore`,
	}
	listEthAccountsCmd = cli.Command{
		Action:    listEthAccounts,
		Name:      "ls",
		Usage:     "list eth accounts under keystore",
		ArgsUsage: "",
		Category:  "LIST ETH ACCOUNT",
		Flags: []cli.Flag{
			KeystoreFlag,
		},
		Description: `list all eth accounts under keystore`,
	}
	getPasswordCmd = cli.Command{
		Action:    getPassword,
		Name:      "getpwd",
		Usage:     "output the real password of private key",
		ArgsUsage: "",
		Category:  "PASSWORD TOOL",
		Flags: []cli.Flag{
			PassFlag,
		},
		Description: `get real password of private key`,
	}
)
