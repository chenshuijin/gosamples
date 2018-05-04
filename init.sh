#!/bin/bash

if [ ! -d ./src/golang.org/x ];then
    mkdir -p src/golang.org/x
fi

git clone https://github.com/golang/text.git src/golang.org/x/text
git clone https://github.com/golang/net.git src/golang.org/x/net
git clone https://github.com/golang/crypto.git src/golang.org/x/crypto
git clone https://github.com/golang/tools.git src/golang.org/x/tools

go install golang.org/x/text/language
go get -u -v gopkg.in/pg.v3
go get -u -v github.com/astaxie/beego
go get -u -v gopkg.in/yaml.v2
go get -u -v gopkg.in/alecthomas/kingpin.v2
go get -u -v gopkg.in/urfave/cli.v1
go get -u -v github.com/hyperledger/fabric
go get -u -v github.com/ethereum/go-ethereum
go get -u -v github.com/gorilla/mux
go get -v -u github.com/nsf/gocode
go get -v -u github.com/rogpeppe/godef
go get -v -u github.com/bradfitz/goimports
go get -v -u github.com/kardianos/govendor
go get -v -u github.com/syndtr/goleveldb/leveldb
