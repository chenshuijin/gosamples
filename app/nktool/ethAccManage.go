package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	cli "gopkg.in/urfave/cli.v1"
)

func newEthAccount(ctx *cli.Context) error {
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(getKeystore(ctx), scryptN, scryptP)
	//pwd := getRealPwd(ctx.String(PassFlag.Name))
	pwd := ctx.String(PassFlag.Name)
	acc, err := ks.NewAccount(pwd)
	if err != nil {
		return nil
	}
	jbs, err := json.Marshal(acc)
	if err != nil {
		log.Println("json marshal err:", err)
	}
	log.Printf("acc:%s\n", jbs)
	return nil
}

func listEthAccounts(ctx *cli.Context) error {
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(getKeystore(ctx), scryptN, scryptP)
	accs := ks.Accounts()
	for idx, acc := range accs {
		jbs, err := json.Marshal(acc)
		if err != nil {
			log.Println("json marshal err:", err)
		}
		log.Printf("acc %d:%s\n", idx, jbs)

	}
	return nil
}

func decEthKey(ctx *cli.Context) error {
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	pwd := ctx.String(PassFlag.Name)
	ks := keystore.NewKeyStore(getKeystore(ctx), scryptN, scryptP)
	kidx := 0
	accs := ks.Accounts()
	log.Println("url:", accs[kidx].URL.Path)

	json, err := ioutil.ReadFile(accs[kidx].URL.Path)
	if err != nil {
		log.Fatal("read key file err:", err)
	}
	key, err := keystore.DecryptKey(json, pwd)
	if err != nil {
		log.Fatal("decrypt key err:", err)
	}

	pribs := crypto.FromECDSA(key.PrivateKey)
	pub := key.PrivateKey.PublicKey
	pubbs := crypto.FromECDSAPub(&pub)
	log.Println("pri:", pribs)
	log.Printf("pri:%x\n", pribs)
	log.Println("pub:", pubbs)
	log.Printf("pub:%x\n", pubbs)
	return nil
}

func getPassword(ctx *cli.Context) error {
	fmt.Println(getRealPwd(ctx.String(PassFlag.Name)))
	return nil
}

func getRealPwd(pwd string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}

func getKeystore(ctx *cli.Context) string {
	kflag := expandPath(ctx.String(KeystoreFlag.Name))
	fullPath, err := filepath.Abs(kflag)
	if err != nil {
		panic(err)
	}
	fmt.Println("fullpath:", fullPath)
	return fullPath
}

// Expands a file path
// 1. replace tilde with users home dir
// 2. expands embedded environment variables
// 3. cleans the path, e.g. /a/b/../c -> /a/c
// Note, it has limitations, e.g. ~someuser/tmp will not be expanded
func expandPath(p string) string {
	if strings.HasPrefix(p, "~/") || strings.HasPrefix(p, "~\\") {
		if home := homeDir(); home != "" {
			p = home + p[1:]
		}
	}
	return path.Clean(os.ExpandEnv(p))
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
