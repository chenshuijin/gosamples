package main

import (
	"encoding/asn1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/go-ray/logging"
	cli "gopkg.in/urfave/cli.v1"
)

func qbal(ctx *cli.Context) error {
	//queryWithKeyJsonBalance()
	//queryEthAccountsBalance()
	queryEthAccBalInStore()
	return nil
}

func queryWithKeyJsonBalance() {
	a := Account{}
	for i := 1; i < 1265; i++ {
		accs := a.GetById(i, i+1)
		logging.Debugf("a:%+v", accs)
		for _, acc := range accs {
			acc.Balance = queryBalance(acc.Address)
			if err := acc.Update(); err != nil {
				logging.Error("udpate acc failed:", err)
			}

		}
	}
}

var w = sync.WaitGroup{}

func queryEthAccBalInStore() {
	i := uint32(0)
	for ; i < 3574; atomic.AddUint32(&i, uint32(1)) {
		for j := 0; j < 10; j++ {
			w.Add(1)
			k := atomic.AddUint32(&i, uint32(1))
			p := path.Join(conf.KeyStore, "key."+strconv.Itoa(int(k)))
			logging.Error("start key file:", k)
			go queryEthAccBalInfile(p)
		}
		w.Wait()
	}
}

func queryEthAccBalInfile(p string) {
	f, err := os.Open(p)
	if err != nil {
		logging.Error("open key file failed:", err)
		return
	}
	defer f.Close()
	defer func() {
		w.Done()
	}()

	b := make([]byte, 128)
	a := &Key2File{}
	off := int64(0)
	for {
		if off%128000 == 0 {
			logging.Error("off set now is:", off)
		}
		_, err = f.ReadAt(b, off)
		if err != nil {
			if err == io.EOF {
				logging.Error("file end")
				break
			}
		}
		off += 128

		_, err = asn1.Unmarshal(b[:], a)
		if err != nil {
			logging.Errorf("unmarshal bytes[%x] to key failed:%+v", b, err)
		}
		addr := fmt.Sprintf("0x%x", a.Address)
		//bal := queryBalance(addr)
		bal := queryBalFromLocal(addr)
		if bal != "0" && bal != "0x0" {
			logging.Errorf("bal:%s,a:%x", bal, a)
			aval := BalAcc{
				Balance: bal,
				Address: addr,
				Public:  a.Public,
				Private: a.Private,
			}
			writeBal2file(aval)
		}
	}
}

func queryEthAccountsBalance() {
	a := EthAccount{}
	for i := 1; i < 1265; i++ {
		accs := a.GetById(i)
		logging.Debugf("a:%+v", accs)
		for _, acc := range accs {
			acc.Balance = queryBalance(acc.Address)
			if err := acc.UpdateOne(); err != nil {
				logging.Error("udpate acc failed:", err)
			}
		}
	}
}

func writeBal2file(a BalAcc) {
	data, err := json.Marshal(a)
	if err != nil {
		logging.Error("marshal balacc failed:", err)
	}
	data = append(data, []byte("\n")...)

	p := path.Join(conf.KeyStore, "vaccs")
	f, err := os.OpenFile(p, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			if e1 := ioutil.WriteFile(p, data, 0644); e1 != nil {
				logging.Error("open vaccs file failed:", err)
			}
		} else {
			logging.Error("open vaccs file failed:", err)
		}
		return
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		logging.Error("write key to file failed:", err)
	}
}

func queryBalance(address string) string {
	resp, err := Get("https://api.etherscan.io/api?module=account&action=balance&tag=latest&apikey=YourApiKeyToken&address=" + address)
	if err != nil {
		logging.Error("post to web api failed:", err)
		return "0"
	}
	logging.Debug("resp:", string(resp))
	br := &BalResp{}
	err = json.Unmarshal(resp, br)
	if err != nil {
		logging.Error("unmarshal resp data failed:", err)
		return "0"
	}

	return br.Result
}

func queryBalFromLocal(addr string) string {
	bq := &BalReq{
		Jsonrpc: "2.0",
		Method:  "eth_getBalance",
		Params:  []string{addr, "latest"},
		Id:      1,
	}
	bdata, err := json.Marshal(bq)
	if err != nil {
		logging.Error("marshal balreq failed:", err)
		return "0x0"
	}
	rdata, err := Post("http://10.60.81.144:8545", bdata)
	if err != nil {
		logging.Errorf("get balance from local node failed, addr[%s]:%v", addr, err)
		return "0x0"
	}

	br := &BalResp{}
	err = json.Unmarshal(rdata, br)
	if err != nil {
		logging.Error("unmarshal resp data failed:", err)
		return "0x0"
	}
	return br.Result
}

type BalReq struct {
	Method  string
	Params  []string
	Id      int
	Jsonrpc string
}

type BalResp struct {
	Jsonrpc string
	Id      int
	Result  string
}

type BalAcc struct {
	Balance string `json:"bal"`
	Address string `json:"addr"`
	Public  []byte `json:"pub"`
	Private []byte `json:"pri"`
}
