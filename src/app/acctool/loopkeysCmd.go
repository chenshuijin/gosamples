package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-ray/logging"
	cli "gopkg.in/urfave/cli.v1"
)

func LoopKeys(ctx *cli.Context) error {
	for {
		a, s := newKey()
		//storeKey2Db(a, s)
		storeKey2file(a, s)
		logging.Debugf("a:%+v", a)
		logging.Debugf("s:%+v", s)
	}
}

func newKey() (*EthAccount, *Secp256Key) {
	privKey, err := NewSecp256Key()
	if err != nil {
		logging.Error("GenerateKey failed:", err)
	}

	privBytes := crypto.FromECDSA(privKey)
	pubKey := privKey.PublicKey
	pubBytes := crypto.FromECDSAPub(&pubKey)
	addr := crypto.PubkeyToAddress(pubKey)

	s := &Secp256Key{
		Public:  pubBytes,
		Private: privBytes,
	}

	a := &EthAccount{
		Address: addr.String(),
		Public:  pubBytes,
	}

	return a, s
}

func storeKey2Db(a *EthAccount, s *Secp256Key) {
	if err := a.Store(); err != nil {
		logging.Errorf("save account[%+v] to db failed:", a, err)
	}
	if err := s.Store(); err != nil {
		logging.Errorf("save secp256 key [%+v] to db failed:", s, err)
	}
}

func storeKey2file(a *EthAccount, s *Secp256Key) {
	keyfilepath := path.Join(conf.KeyStore, a.Address)
	keyfilepath, _ = filepath.Abs(keyfilepath)

	d := make([]byte, 120)
	d[0] = byte(len(a.Address))
	d = append(d[:], []byte(a.Address)[:]...)

	d[len(d)+1] = byte(len(s.Public))
	d = append(d[:], s.Public[:]...)

	d[len(d)+1] = byte(len(s.Private))
	d = append(d[:], s.Private[:]...)

	writeKey2file(d)
}

func getstorefilepath() string {
	index := path.Join(conf.KeyStore, "index")
	index, _ = filepath.Abs(index)
	currentfilebase := path.Join(conf.KeyStore, "key.")
	currentfile := currentfilebase + getIndex(false)

	f, err := os.Open(currentfile)
	if err != nil {
		os.Create(currentfile)
		f, _ = os.Open(currentfile)
	}
	defer f.Close()

	finfo, _ := f.Stat()
	if finfo.Size() > 1024*100 {
		currentfile = currentfilebase + getIndex(true)
	}

	return currentfile
}

func getIndex(update bool) string {
	index := path.Join(conf.KeyStore, "index")
	index, _ = filepath.Abs(index)
	data, e1 := ioutil.ReadFile(index)
	if e1 != nil {
		if e1 := ioutil.WriteFile(index, []byte("0"), 0644); e1 != nil {
			logging.Error("write index failed:", e1)
		}
		return "0"
	} else {
		i, e2 := strconv.ParseUint(string(data), 10, 64)
		if e2 != nil {
			return "0"
		} else {
			if update {
				if e3 := ioutil.WriteFile(index, []byte(string(i+1)), 0644); e3 != nil {
					logging.Error("write index failed:", e3)
				}
				return string(i + 1)
			} else {
				return string(i)
			}
		}
	}
}

func writeKey2file(data []byte) {
	cpath := getstorefilepath()
	if err := ioutil.WriteFile(cpath, data, os.ModeAppend); err != nil {
		logging.Error("write key to file failed:", err)
	}
}
