use std::env;
use std::collections::HashMap;

fn main() {
    env::set_var("RUST_BACKTRACE", "1");

    let mut reviews: HashMap<String, String> = HashMap::new();

    reviews.insert(String::from("Ancient Roman History"), String::from("Very accurate."));
    reviews.insert(String::from("Cooking with Rhubarb"), String::from("Sweet recipes."));
    reviews.insert(String::from("Programming in Rust"), String::from("Great examples."));

    // Look for a specific review
    let book: &str = "Programming in Rust";
    println!("\nReview for \'{}\': {:?}", book, reviews.get(book));

    let obsolete: &str = "Ancient Roman History";
    println!("\n'{}\' removed.", obsolete);
    reviews.remove(obsolete);

    println!("\nReview for \'{}\': {:?}", obsolete, reviews.get(obsolete));

    let mut counter = 1;
    let stop_loop = loop {
        counter *= 2;
        if counter > 10 {
            break counter;
        }
    };
    println!("Break the loop at counter = {}.", stop_loop);

    while counter < 50 {
        println!("We loop the while...");
        counter += 10;
    }

    let big_birds = ["ostrich", "peacock", "stork"];
    for bird in big_birds.iter() {
        println!("The {} is a big bird.", bird);
    }

    for number in 0..5 {
        println!("{}", number * 2);
    }

    let v = vec![1, 2, 3, 4, 5];
    // println!("{}", v[6]);
    println!("{}", v[2]);

    let fruits = vec!["banana", "apple", "coconut", "orange", "strawberry"];
    let first = fruits.get(0);
    println!("{:?}", first);
    let non_existent = fruits.get(101);
    println!("{:?}", non_existent);

    for &index in [0, 2, 101].iter() {
        match fruits.get(index) {
            Some(&"coconut") => println!("Coconuts are awesome!!!"),
            Some(fruit_name) => println!("It's a delicious {}!", fruit_name),
            None => println!("There is no fruit! :("),
        }
    }

    let a_number: Option<u8> = Some(7);
    match a_number {
        Some(7) => println!("That's my lucky number!"),
        _ => {},
    }

    let gift = Some("candy");
    assert_eq!(gift.unwrap(), "candy");
    // let empty_gift: Option<&str> = None;
    // assert_eq!(empty_gift.unwrap(), "candy"); // This will panic!

    let a = Some("value");
    assert_eq!(a.expect("fruits are healthy"), "value");

    // let b: Option<&str> = None;
    // b.expect("fruits are healthy"); // panics with `fruits are healthy`

    assert_eq!(Some("dog").unwrap_or("cat"), "dog");
    assert_eq!(None.unwrap_or("cat"), "cat");
}
