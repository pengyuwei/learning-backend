# Rust

## Catalogue

- ProgrammingRust: 《Programming Rust》的学习；
- trpl: 《The Rust Programming Language》（RUST权威指南）的学习

## Environment

Cargo 是 Rust 工具链中内置的构建系统和包管理器。

```
sudo apt install cargo
cargo new hello 
cd hello
cargo check
cargo add curl
cargo run main.rs
cargo build
cargo doc
cargo build --release
cargo update
```

依赖组件需要写到Cargo.toml里。版本查询：https://crates.io/

## 更换国内源

~/.cargo/config

```
wget http://www.memcd.com/conf/rust/cargo/config
```

## Reference

- https://docs.rs/chrono/0.4.20/chrono/struct.DateTime.html#method.format
- https://crates.io/
- https://kaisery.github.io/trpl-zh-cn/ch14-01-release-profiles.html
