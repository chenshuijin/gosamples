package main

import (
	"bytes"
	"encoding/asn1"
	"encoding/hex"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"path/filepath"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-ray/logging"
	cli "gopkg.in/urfave/cli.v1"
)

func LoopKeys(ctx *cli.Context) error {
	for {
		a, s := newKey()
		//storeKey2Db(a, s)
		storeKey2file(a, s)
		logging.Debugf("a:%+v", a.Address)
		logging.Debugf("s.pub:%+x", s.Public)
		logging.Debugf("s.priv:%+x", s.Private)
	}
	return nil
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

	addr, err := hex.DecodeString(a.Address[2:])
	if err != nil {
		logging.Error("decode address to bytes failed:", err)
	}

	k2f := Key2File{
		Address: addr,
		Public:  s.Public,
		Private: s.Private,
	}

	bs, err := asn1.Marshal(k2f)
	if err != nil {
		logging.Error("asn1 marshal failed:", err)
	}
	logging.Debugf("k2f:%+x", bs)

	pad := bytes.Repeat([]byte{0x00}, 128-len(bs))
	bs = append(bs, pad...)
	writeKey2file(bs)
}

func getstorefilepath() string {
	index := path.Join(conf.KeyStore, "index")
	index, _ = filepath.Abs(index)
	currentfilebase := path.Join(conf.KeyStore, "key.")
	logging.Debug("get index:", getIndex(false))
	currentfile := currentfilebase + getIndex(false)
	logging.Debug("currentfile:", currentfile)
	f, err := os.Open(currentfile)
	if err != nil {
		logging.Error("open current key file failed:", err)
		os.Create(currentfile)
		f, _ = os.Open(currentfile)
	}
	defer f.Close()

	finfo, e1 := f.Stat()
	if e1 != nil {
		logging.Error("get key file stat failed:", e1)
	}
	//if finfo.Size() > 1024*1024*100 {
	if finfo.Size() > conf.PerKeyFileLength {
		currentfile = currentfilebase + getIndex(true)
	}

	return currentfile
}

func getIndex(update bool) string {
	index := path.Join(conf.KeyStore, "index")
	index, _ = filepath.Abs(index)
	data, e1 := ioutil.ReadFile(index)
	if e1 != nil {
		logging.Error("read index file failed:", e1)
		if e2 := os.Mkdir(conf.KeyStore, os.ModePerm); e2 != nil {
			logging.Error("mkdir failed:", e2)
		}
		i := big.NewInt(0)
		if e1 := ioutil.WriteFile(index, i.Bytes(), 0644); e1 != nil {
			logging.Error("write index failed1 :", e1)
		}
		return i.String()
	} else {
		i := &big.Int{}
		i = i.SetBytes(data)

		if update {
			b1 := big.NewInt(1)
			i = i.Add(i, b1)
			if e3 := ioutil.WriteFile(index, i.Bytes(), 0644); e3 != nil {
				logging.Error("write index failed2 :", e3)
			}
			return i.String()
		} else {
			return i.String()
		}
	}
}

func writeKey2file(data []byte) {
	cpath := getstorefilepath()
	f, err := os.OpenFile(cpath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			if e1 := ioutil.WriteFile(cpath, data, 0644); e1 != nil {
				logging.Error("open key file failed:", err)
			}
		} else {
			logging.Error("open key file failed:", err)
		}
		return
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		logging.Error("write key to file failed:", err)
	}
}
