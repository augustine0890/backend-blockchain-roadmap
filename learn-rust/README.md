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

