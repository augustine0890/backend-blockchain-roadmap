fn main() {
    let mut language = "Rust";
    println!("Language: {}", language);
    language = "Golang";
    println!("Language: {}", language);
    let (course, category) = ("Rust", "beginer");
    println!("This is a {} course in {}", category, course);
}
