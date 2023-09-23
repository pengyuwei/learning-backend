.DEFAULT_GOAL := all
.PHONY : all
all : build

init:
	go mod init mysql-example
	go get -u golang.org/x/lint
	go get github.com/go-sql-driver/mysql
	go get gorm.io/driver/mysql
	go get gorm.io/gorm
build:
	go build main.go
clean:
	@rm main