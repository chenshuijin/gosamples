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
OLD_GOPATH:=$(GOPATH)
export GOPATH=$(CURDIR):$(OLD_GOPATH)

all: clean test app
quick: app

app:
	echo $(GOPATH)
	go fmt app/...
	go install app/...
test:
	go test -v -bench=. app/...
	go test -v -bench=. -cpuprofile=cpu.profile app/tunnel
clean:
	rm -rf $(CURDIR)/bin/
goget:
	go get -u -v golang.org/x/text/language
	go get -u -v gopkg.in/pg.v3
