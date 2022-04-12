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

# Rustlang
- memory: `u8` is 8 bits (1 byte), u16 is 16 bits (2 bytes), ...