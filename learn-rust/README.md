# Install Rust
- Install [Rust](https://www.rust-lang.org/tools/install)
  - `rustc --version` verify installing
- Config the PATH environment variable
  - `export PATH="$HOME/.cargo/bin:$PATH"`
  - `echo "export PATH='\$HOME/.cargo/bin:\$PATH'" >> ~/.zshrc`
  - `source ~/.zshrc`
  - `export PATH=$HOME/bin:/usr/local/bin:$PATH` then
  - `source ~/.zshrc`

# Project
- Create a new project
  - `cargo new <name-project>`
- Compile: `cargo build`
- Compile and run: `cargo run`
- Remove generated `target`
  - `cargo clean`
- `Cargo` manager has all the rust libraries and frameworks.
# Rustlang
- Rust does not have a garbage collector --> ownership and its borrowing.
- Rust can interoperate with C language. It provides a foreign function interface to communicate with C API's.
- Data race is a condition where two or more threads can acces the same memory location.
  - Rust uses the concept of ownership to avoid data races.
- memory: `u8` is 8 bits (1 byte), u16 is 16 bits (2 bytes), ...