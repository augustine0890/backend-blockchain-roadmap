fn main() {
    println!("Hello, world!");
    println!("Hello");
    println!("Number: {}", 1);
    println!("{} is a {} company", "Supertree", "Game Blockchain");
    println!("{0} is a {1} company", "Supertree", "Game Blockchain");
    println!("{company} is a {kind} company", company = "Supertree", kind = "Game Blockchain");
    println!("Number: 10\nBinary:{:b} Hexadecimal:{:x} Octal:{:o}", 10, 10, 10);
    println!("{:?}", ("This is a Rust Course", 101));
    eprintln!("Rustlang");
    println!("{:?}", ("This is a Rust Course", 101));

    test();
}

fn test() {
    println!("I am learning Rust programming language");
    print!("{}\n{}\n{}\n{}\n{}\n", 1, 22, 333, 4444, 55555);
}
