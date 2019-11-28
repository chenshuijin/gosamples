package main

import cli "gopkg.in/urfave/cli.v1"

var (
	newEthAccountCmd = cli.Command{
		Action:    newEthAccount,
		Name:      "newe",
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
	decEthKeyCmd = cli.Command{
		Action:    decEthKey,
		Name:      "dec",
		Usage:     "decrypt eth key under keystore",
		ArgsUsage: "",
		Category:  "DECRYPT ETH KEY",
		Flags: []cli.Flag{
			KeystoreFlag,
			PassFlag,
		},
		Description: `decrypt eth key under keystore`,
	}
	newNasAccountCmd = cli.Command{
		Action:    newNasAccount,
		Name:      "newn",
		Usage:     "create new nas account under keystore",
		ArgsUsage: "",
		Category:  "NEW NAS ACCOUNT",
		Flags: []cli.Flag{
			PassFlag,
			KeystoreFlag,
		},
		Description: `create new nas account under keystore`,
	}
	eth2nasKeyCmd = cli.Command{
		Action:    eth2nasKey,
		Name:      "e2nk",
		Usage:     "convert eth account to nas account under keystore",
		ArgsUsage: "",
		Category:  "NAS ACCOUNT",
		Flags: []cli.Flag{
			PassFlag,
			KeystoreFlag,
		},
		Description: `convert eth account to nas account under keystore`,
	}
	pub2addrCmd = cli.Command{
		Action:      pub2nasaddr,
		Name:        "nasaddr",
		Usage:       "convert public key to nas address",
		ArgsUsage:   "",
		Category:    "NAS TOOL",
		Flags:       []cli.Flag{},
		Description: `convert public key to nas address`,
	}
)
