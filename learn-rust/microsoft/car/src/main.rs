// Declare Car struct to describe vehicle with four named fields
#[derive(PartialEq, Debug)]
struct Car {
    color: String,
    motor: Transmission,
    roof: bool,
    age: (Age, u32)
}

// Declare enum for Car transmission type
#[derive(PartialEq, Debug)]
enum Transmission {
    Manual,
    SemiAuto,
    Automatic
}

#[allow(dead_code)]
#[derive(PartialEq, Debug)]
enum Age {
    New,
    Used
}

fn car_factory(color: String, motor: Transmission, roof: bool, miles: u32) -> Car {
    // Create a new "Car" instance as requested
    // - Bind first three fields to values of input arguments
    // Corrected code: "mileage" is replaced with "age"
    // Corrected code: Bind "age" to tuple returned from car_quality(miles)
    Car {
        color: color,
        motor: motor,
        roof: roof,
        age: car_quality(miles)
    }
}

fn car_quality(miles: u32) -> (Age, u32) {
    // Declare and initialize the return tuple value
    // For a new car, set the miles to 0
    // Corrected code: Define "quality"
    // - Set the value to a "New" car
    // - Set the mileage using the "miles" input argument
    let quality: (Age, u32) = (Age::New, miles);
    quality
}

fn main() {
    // Create car color array
    
    let colors = ["Blue", "Green", "Red", "Silver"];

    // Declare the car type and initial values
    let mut car: Car;
    let mut engine: Transmission = Transmission::Manual;

    // Order 3 cars, one car for each type of transmission

    // Car order #1: New, Manual, Hard top
    car = car_factory(String::from(colors[0]), engine, true, 0);
    println!("Car order 1: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    // Car order #2: Used, Semi-automatic, Convertible
    engine = Transmission::SemiAuto;
    car = car_factory(String::from(colors[1]), engine, false, 100);
    println!("Car order 2: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    // Car order #3: Used, Automatic, Hard top
    engine = Transmission::Automatic;
    car = car_factory(String::from(colors[2]), engine, true, 200);
    println!("Car order 3: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);
}
