.DEFAULT_GOAL := all

.PHONY : all
all : default

default:
	mkdir -p bin
	gcc -Wall -g -o bin/c.debug hello.c
	gcc -O1 -o bin/c.O1 hello.c
	gcc -O2 -o bin/c.O2 hello.c
	gcc -O4 -o bin/c.O4 hello.c
	g++ -Wall -g -o bin/cpp.debug hello.cpp
	g++ -O1 -o bin/cpp.O1 hello.cpp
	g++ -O2 -o bin/cpp.O2 hello.cpp
	g++ -O4 -o bin/cpp.O4 hello.cpp
	cargo build
	cp -f target/debug/rust bin/rust.debug
	cargo build --release
	cp -f target/release/rust bin/rust.release
	nim c -d:release -o:bin/nim -r hello.nim
	zig build-exe -O ReleaseSmall --name zig hello.zig
	mv ./zig bin/
	go build -o bin/go hello.go

	# ocamlopt.opt -compact -o bin/ocaml hello.ml
	# fpc -Os hello.pas
	# ghc hello.hs
	# gdc -Os -o hello hello.d

	ls -lh bin

	ls -lh bin

clean:
	@rm -f Cargo.lock
	@rm -rf target
	@rm -rf bin