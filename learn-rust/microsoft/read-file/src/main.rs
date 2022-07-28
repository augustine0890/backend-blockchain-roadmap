use std::fs::File;
use std::io::{Error, Read};
use std::path::PathBuf;

fn read_file_contents(path: PathBuf) -> Result<String, Error> {
    let mut string = String::new();

    // Access a file at a specified path
    // ---------------------------------
    // TODO #1:
    // - Pass variable to `file` variable on success, or
    // - Return from function early if there's an error
    let mut file: File = match File::open(path) {
        Ok(file_handle) => file_handle,
        Err(io_error) => return Err(io_error),
    };

    // Read file contents into `String` variable with `read_to_string`
    // ---------------------------------
    // Success path is already filled in
    // TODO #2: Return from function early if there's an error
    match file.read_to_string(&mut string) {
        Ok(_) => (),
        Err(io_error) => return Err(io_error)
    };

    // TODO #3: Return `string` variable as expected by function signature
    Ok(string)
}

fn copy_and_return<'a>(vector: &'a mut Vec<String>, value: &'a str) -> &'a String {
    vector.push(String::from(value));
    vector.get(vector.len() - 1).unwrap()
}

fn main() {
    if read_file_contents(PathBuf::from("src/main.rs")).is_ok() {
        println!("The program found the main file.");
    }
    if read_file_contents(PathBuf::from("non-existent-file.txt")).is_err() {
        println!("The program reported an error for the file that doesn't exist.");
    }

    let name1 = "Joe";
    let name2 = "Chris";
    let name3 = "Anne";

    let mut names = Vec::new();
    assert_eq!("Joe", copy_and_return(&mut names, &name1));
    assert_eq!("Chris", copy_and_return(&mut names, &name2));
    assert_eq!("Anne", copy_and_return(&mut names, &name3));

    assert_eq!(
        names,
        vec!["Joe".to_string(), "Chris".to_string(), "Anne".to_string()]
    )
}