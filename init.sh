#!/bin/bash

if [ ! -d ./src/golang.org/x ];then
    mkdir -p src/golang.org/x
fi

git clone https://github.com/golang/text.git
git clone https://github.com/golang/net.git
git clone https://github.com/golang/crypto.git
git clone https://github.com/golang/tools.git

go install golang.org/x/text/language
go get -u -v gopkg.in/pg.v3
go get -u -v github.com/astaxie/beego
go get -u -v gopkg.in/yaml.v2
go get -u -v gopkg.in/alecthomas/kingpin.v2
go get -u -v gopkg.in/urfave/cli.v1
go get -u -v github.com/hyperledger/fabric
go get -u -v github.com/ethereum/go-ethereum
go get -u -v github.com/gorilla/mux
