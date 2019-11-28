#!/bin/bash

if [ ! -d ./src/golang.org/x ];then
    mkdir -p src/golang.org/x
fi

git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto
git clone https://github.com/golang/tools.git $GOPATH/src/golang.org/x/tools

go install golang.org/x/text/language

pkgs=(
 gopkg.in/pg.v3
 github.com/astaxie/beego
 gopkg.in/yaml.v2
 gopkg.in/alecthomas/kingpin.v2
 gopkg.in/urfave/cli.v1
 github.com/hyperledger/fabric
 github.com/ethereum/go-ethereum
 github.com/gorilla/mux
 github.com/nsf/gocode
 github.com/rogpeppe/godef
 github.com/bradfitz/goimports
 github.com/kardianos/govendor
 github.com/syndtr/goleveldb/leveldb
 github.com/rs/cors
)

for pkg in ${pkgs[*]}; do
    go get -u -v $pkg
done
