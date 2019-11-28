package main

import (
	"flag"
	"fmt"
	"net"
)

var host = flag.String("host", "localhost:80", "the host of the remote server")
var data = flag.String("d", "", "the data to send")

func main() {
	flag.Parse()
	fmt.Println("the host:", *host)
	addr, err := net.ResolveTCPAddr("tcp", *host)
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.SetNoDelay(true)
	*data = `POST localhost:7050/chaincode\r\n{ "jsonrpc": "2.0", "method": "deploy", "params": { "type": 1, "chaincodeID":{"path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02" }, "CtorMsg": {"function":"init","args":["a", "1000", "b", "2000"] }}, "id": "1"  }`

	count, err := conn.Write([]byte(*data))
	if err != nil {
		panic(err)
	}
	fmt.Println("send count:", count)
	recvData := make([]byte, 2048)
	count, err = conn.Read(recvData)
	if err != nil {
		panic(err)
	}
	fmt.Println("recv count:", count)
	fmt.Println("recv str:", string(recvData))
}
