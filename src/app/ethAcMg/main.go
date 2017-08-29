package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore("keystore", scryptN, scryptP)
	accs := ks.Accounts()
	log.Println("accs:", accs)
	ks.NewAccount("")
	ptks := keystore.NewPlaintextKeyStore("ptkeystore")
	ptks.NewAccount("")
}
