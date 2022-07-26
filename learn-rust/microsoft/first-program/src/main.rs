fn goodbye(message: &str) {
    println!("\n{}", message);
}

fn divide_by_5(num: u32) -> u32 {
    if num == 0 {
        return 0;
    }
    num / 5
}

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

    // Declare a tuple of length 3
    let tuple_e = ('E', 5i32, true);
    println!("Is '{}' the {}th letter of the alphabet? {}", tuple_e.0, tuple_e.1, tuple_e.2);

    // Classic struct with named fields
    struct Student {
        name: String,
        level: u8,
        remote: bool
    }
    // Tuple struct with data types only
    struct Grades(char, char, char, char, f32);

    let user_1 = Student{
        name: String::from("Constance Sharma"),
        remote: true,
        level: 2
    };
    let user_2 = Student {
        name: String::from("Dyson Tan"),
        level: 5,
        remote: false
    };
    let mark_1 = Grades('A', 'A', 'B', 'A', 3.75);
    let mark_2 = Grades('B', 'A', 'A', 'C', 3.25);

    println!("{}, level {}. Remote: {}. Grades: {}, {}, {}, {}. Average: {}",
            user_1.name, user_1.level, user_1.remote, mark_1.0, mark_1.1, mark_1.2, mark_1.3, mark_1.4);
    println!("{}, level {}. Remote: {}. Grades: {}, {}, {}, {}. Average: {}",
            user_2.name, user_2.level, user_2.remote, mark_2.0, mark_2.1, mark_2.2, mark_2.3, mark_2.4);


    // Define a tuple struct
    #[derive(Debug)]
    struct KeyPress(String, char);

    // Define a classic struct
    #[derive(Debug)]
    struct MouseClick{ x: i64, y: i64 }

    // Define the WebEvent enum variants to use the data from the structs
    // and a boolean type for the page Load variant
    #[derive(Debug)]
    enum WebEvent {
        WELoad(bool),
        WEClick(MouseClick),
        WEKeys(KeyPress)
    }

    // Instantiate a KeyPress tuple and bind the key values
    let click = MouseClick{x: 100, y: 250 };
    println!("Mouse click location: {}, {}", click.x, click.y);

    // Instantiate a KeyPress tuple and bind the key values
    let keys = KeyPress(String::from("Ctrl+"), 'N');
    println!("\nKeys pressed: {}{}", keys.0, keys.1);

    // Instantiate WebEvent enum variants
    // Set the boolean page Load value to true
    let we_load = WebEvent::WELoad(true);
    let we_click = WebEvent::WEClick(click);
    let we_key = WebEvent::WEKeys(keys);

    // Use the {:#?} syntax to display the enum structure and data in a readable form
    println!("\nWebEvent enum structure: \n\n {:#?} \n\n {:#?} \n\n {:#?}", we_load, we_click, we_key);

    let formal = "Formal: Good bye.";
    let casual = "Casual: See you later!";
    goodbye(formal);
    goodbye(casual);

    let num = 25;
    println!("{} divided by 5 = {}", num, divide_by_5(num));
}
