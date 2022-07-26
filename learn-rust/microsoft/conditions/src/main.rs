fn main() {
    // Days of the week
    let days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"];

    let first = days[0];
    println!("The first day of the week: {}", first);
    let last = days[days.len() - 1];
    println!("The last day of the week: {}", last);

    // Declare array, initialize all values to 0, length = 5
    let bytes = [0; 5];
    println!("Array: {:?}", bytes);

    // Declare vector, initialize with three values
    let three_nums = vec![15, 3, 46];
    println!("Initial vector: {:?}", three_nums);
    let zeroes = vec![0; 5];
    println!("Zeroes vector: {:?}", zeroes);

    // Create empty vector
    let mut fruit = Vec::new();
    fruit.push("Apple");
    fruit.push("Banana");
    fruit.push("Cherry");
    println!("Fruits: {:?}", fruit);
    println!("Pop off: {:?}", fruit.pop());
    println!("Fruits: {:?}", fruit);

    // Declare vector, initialize with three values
    let mut index_vec = vec![15, 3, 46];
    let three = index_vec[1];
    println!("Vector: {:?}, three = {}", index_vec, three);
    index_vec[1] = index_vec[1] + 5;
    println!("Vector: {:?}", index_vec);

}
