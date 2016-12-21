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

export GOPATH=$(CURDIR)

all: clean test app

app:
	go fmt app/...
	go install app/...
test:
	go test -v app/...
	go test -v bcd
clean:
	rm -rf $(CURDIR)/bin/
goget:
	go get -u -v golang.org/x/text/language
	go get -u -v gopkg.in/pg.v3
