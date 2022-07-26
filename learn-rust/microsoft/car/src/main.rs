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

// #[allow(dead_code)]
#[derive(PartialEq, Debug)]
enum Age {
    New,
    Used
}

fn car_factory(color: String, motor: Transmission, roof: bool, miles: u32) -> Car {
    // Show details about car order
    // Corrected code: If order is for Used car, check roof type, print details
    // Corrected code: Else, order is for New car, check roof tye, print details
    // Call the `println!` macro to show the car order details
    if car_quality(miles).0 == Age::Used {
        if roof {
            println!("Preparing a used car: {:?}, {}, Hard top, {} miles", motor, color, miles);
        } else {
            println!("Preparing a used car: {:?}, {}, Convertible, {} miles", motor, color, miles);
        }
    } else {
        if roof {
            println!("Building a new car: {:?}, {}, Hard top, {} miles", motor, color, miles);
        } else {
            println!("Building a new car: {:?}, {}, Convertible, {} miles", motor, color, miles);
        }
    }
    Car {
        color: color,
        motor: motor,
        roof: roof,
        age: car_quality(miles)
    }
}

fn car_quality(miles: u32) -> (Age, u32) {
    if miles > 0 {
        return (Age::Used, miles);
    }
    (Age::New, miles)
}

fn main() {
    // Create car color array
    
    // let colors = ["Blue", "Green", "Red", "Silver"];

    // // Declare the car type and initial values
    // let mut car: Car;
    // let mut engine: Transmission = Transmission::Manual;

    // // Order 3 cars, one car for each type of transmission

    // // Car order #1: New, Manual, Hard top
    // car = car_factory(String::from(colors[0]), engine, true, 0);
    // println!("Car order 1: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    // // Car order #2: Used, Semi-automatic, Convertible
    // engine = Transmission::SemiAuto;
    // car = car_factory(String::from(colors[1]), engine, false, 100);
    // println!("Car order 2: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    // // Car order #3: Used, Automatic, Hard top
    // engine = Transmission::Automatic;
    // car = car_factory(String::from(colors[2]), engine, true, 200);
    // println!("Car order 3: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    // Car order #1: New, Manual, Hard top
    car_factory(String::from("Orange"), Transmission::Manual, true, 0);
    
    // Car order #2: Used, Semi-automatic, Convertible
    car_factory(String::from("Red"), Transmission::SemiAuto, false, 565);

    // Car order #3: Used, Automatic, Hard top
    car_factory(String::from("White"), Transmission::Automatic, true, 3000);
}
