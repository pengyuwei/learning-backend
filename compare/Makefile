.DEFAULT_GOAL := all

.PHONY : all
all : default

default:
	mkdir -p bin
	gcc -O2 -o bin/c hello.c
	g++ -O2 -o bin/cpp hello.cpp
	cargo build --release
	cp -f target/release/rust bin/rust
	nim c -d:release -o:bin/nim -r hello.nim
	zig build-exe -O ReleaseSmall --name zig hello.zig
	mv ./zig bin/
	go build -o bin/go hello.go

	# ocamlopt.opt -compact -o bin/ocaml hello.ml
	# fpc -Os hello.pas
	# ghc hello.hs
	# gdc -Os -o hello hello.d

	ls -lh bin

clean:
	@rm -f Cargo.lock
	@rm -rf target
	@rm -rf bin