package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	cli "gopkg.in/urfave/cli.v1"
)

func unlockAccount(ctx *cli.Context) error {
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(getKeystore(ctx), scryptN, scryptP)
	accs := ks.Accounts()
	log.Println("accs:", accs)
	err := ks.Unlock(accs[0], ctx.String(PassFlag.Name))
	if err != nil {
		log.Println("err:", err)
	}
	return nil
}
