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
- Compile: `cargo build` --> build a project.
- Compile and run: `cargo run` --> build and run.
- Check project types: `cargo check`
- Remove generated `target`
  - `cargo clean`
- Build documentation: `cargo doc`.
- Publish a library to crates.io: `cargo publish`.
- `Cargo` manager has all the rust libraries and frameworks.
# Rustlang
- Rust is a statically typed language. The compiler must know the exact data type for all variables in your code for your program to compile and run.
- Rust does not have a garbage collector --> ownership and its borrowing.
- Rust can interoperate with C language. It provides a foreign function interface to communicate with C API's.
- Data race is a condition where two or more threads can acces the same memory location.
  - Rust uses the concept of ownership to avoid data races.
- memory: `u8` is 8 bits (1 byte), u16 is 16 bits (2 bytes), ...
- Rust supports text values with two basic string types and one character type. A character is a single item, while a string is a series of characters. All text types are valid UTF-8 representations.
- Integers are identified by a combination of its bit size and whether it's signed (`i`) or unsigned (`u`).
- All text types are valid UTF-8 representations.
- A tuple is a collection of values of different types. The data type is based on the data types of its elements, and the length is fixed based on the number of elements.
- String data that's stored inside another data structure, such as a struct or vector, must be converted from a string literal reference(`&str`) to a `String` type.

## Arrays
- An array is a collection of objects of the same type that are stored sequentially in memory.
## Vector
- Vectors can store multiple values that have the same data type.
- The size or length of a vector can grow or shrink at any time.
## Functions in Rust
- Every Rust program must have one function named `main`. The code in the `main` function is always the first code run in a Rust program.
- We can call other functions from within the `main` function, or from within other functions.
