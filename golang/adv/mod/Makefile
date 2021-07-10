.DEFAULT_GOAL := all

.PHONY : all
all : linuxst
linux: linuxdy

linuxdy : main.go lib.go
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o demo.out main.go lib.go

linuxst: main.go lib.go
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o demo.out main.go lib.go

mac:
	set CGO_ENABLED=0
	set GOOS=darwin3
	set GOARCH=amd64
	go build -o demo.out main.go lib.go

clean:
	@rm -f *.out
	@rm -f demo