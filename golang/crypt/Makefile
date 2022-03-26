.DEFAULT_GOAL := all
.PHONY : all
all : linux
linux: build

GOFLAGS = -gcflags "-N -l"

build : *.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o crypt crypt.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o example example.go

mac:
	mkdir -p bin
	set CGO_ENABLED=0
	set GOOS=darwin3
	set GOARCH=amd64
	go build $(GOFLAGS) -o crypt crypt.go
	go build $(GOFLAGS) -o example example.go

clean:
	rm -f crypt
	rm -f *.out
	rm -f bin/*