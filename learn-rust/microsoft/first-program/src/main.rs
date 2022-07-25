fn main() {
    // Declare a variable
    let mut a_number = 10;
    println!("The number is {}", a_number);

    // Declare a second variable and bind the value
    let a_word = "Ten";

    // Bind a value to the first variable
    a_number = 15;

    println!("The number is {}", a_number);
    println!("The word is {}.", a_word);

    let shadow_number = 5;

    let shadow_number = shadow_number + 5;

    let shadow_number = shadow_number * 2;
    
    println!("The shadow number is {}", shadow_number);

    let number: u32 = 14;
    println!("The number is {}", number);

    let number_64 = 4.0; // compiler infers the value to use the default type f64
    println!("The f64 number is {}", number_64);

    let number_32: f32 = 4.5; // type f32 specified via annotation
    println!("The f32 is {}", number_32);

    // Addition, Subtraction, and Multiplication
    println!("1 + 2 = {} and 8 - 5 = {} and 15 * 3 = {}", 1u32 + 2, 8 - 5, 15 * 3);
    println!("9 / 2 = {} but 9.0 / 2.0 = {}", 9u32 / 2, 9.0 / 2.0);

    // Declare variable to store result of "greater than" test, Is 1 > 4? --false
    let is_bigger = 1 > 4;
    println!("Is 1 > 4? {}", is_bigger);

    let uppercase_s: char = 'S';
    println!("Uppercase s is {}", uppercase_s);
    let lowercase_f: char = 'f';
    println!("Lowercase f is {}", lowercase_f);
    let smiley_face = '😃';
    println!("Smiley face: {}", smiley_face);
    // Compiler interprets a series of items in quotations as a "str" data type and creates a "&str" reference
    let string_1 = "miley ";
    println!("miley {}", string_1);
    // Specify the data type "str" with the reference syntex "&str"
    let string_2: &str = "ace";
    println!("string 2: {}", string_2);
}
