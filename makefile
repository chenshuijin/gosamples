# To make samples

# OS Detection

ifeq ($(OS),Windows_NT)
	# Left here for when/if we will support building on windows
	IS_WINDOWS:=true
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		IS_LINUX:=true
	endif
	ifeq ($(UNAME_S),Darwin) # Mac OS X
		IS_MAC_OS_X:=true
	endif
endif

all: clean test build
#quick: app

build:
	go fmt ./app/...
	mkdir -p $(CURDIR)/bin
	go build -o $(CURDIR)/bin/ ./app/...
test:
	go test -v -bench=. ./app/...
	go test -v -bench=. -cpuprofile=cpu.profile ./app/tunnel
clean:
	rm -rf $(CURDIR)/bin/
goget:
	go install golang.org/x/text/language
	go get -u -v gopkg.in/pg.v3
	go get -u -v github.com/astaxie/beego
	go get -u -v gopkg.in/yaml.v2
	go get -u -v gopkg.in/alecthomas/kingpin.v2
