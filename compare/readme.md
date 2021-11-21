# 语言对比

- c: C是一种通用的编程语言，广泛用于系统软件与应用软件的开发。于1969年至1973年间诞生。
- cpp: C++是一种被广泛使用的计算机程序设计语言。
- zig: Zig是一种通用编程语言和工具链，用于维护稳健、优化和可重用的软件。
- rust: 一种赋予每个人力量的语言,建立可靠和高效的软件。
- nim: Nim是一种静态类型的编译系统编程语言。
- go: Go是一种开源编程语言，可以轻松构建简单、可靠和高效的软件。
- ocamlopt: Native-code compilation
- fpc: Free Pascal Compiler
- ghc: Glasgow Haskell Compiler
- gdc: implementation of programming language D

## 可执行文件大小对比

不同语言的HelloWorld文件大小对比

```
17K 	 c
17K 	 cpp
78K 	 zig
91K 	 nim
1.7M 	 go
9.6M 	 rust
```

环境

```
uname -a
Linux ubuntu2-20210214-1307 5.11.0-1016-oracle #17~20.04.1-Ubuntu SMP Thu Aug 12 06:20:42 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux

sudo apt install build-essential
sudo apt install nim

wget https://ziglang.org/builds/zig-linux-x86_64-0.9.0-dev.1678+7747bf07c.tar.xz
tar -xvf zig-linux-x86_64-0.9.0-dev.1678+7747bf07c.tar.xz
sudo mv zig-linux-x86_64-0.9.0-dev.1678+7747bf07c /usr/local/zig-linux-x86_64-0.9.0-dev.1678
sudo ln -s /usr/local/zig-linux-x86_64-0.9.0-dev.1678/zig /usr/local/bin/zig

wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz
tar zxvf go1.17.3.linux-amd64.tar.gz
sudo mv go /usr/local/
sudo ln -s /usr/local/go/bin/go /usr/local/bin/go

```