var userID;

userID = 12;
console.log(typeof userID);
userID = 'user1'
console.log(typeof userID);

const car = {
  wheels: 4,
  color: "red",
  drive: function(){
    console.log("wroom wroom")
  }
}
console.log(Object.keys(car));
console.log(Object.keys(car)[0]);
console.log(typeof Object.keys(car)[0]);
car.drive();

const cars = {
  ferrari: "california",
  porsche: "911",
  bugatti: "veyron",
}
const key = "ferrari"
console.log(cars.key);
console.log(cars[key]);

let secondCar = car;
car.color = "bule";
console.log(car);
console.log(secondCar);

const thirdCar = Object.assign({}, car);
car.wheels = 8;
console.log(car);
console.log(thirdCar);

const fruitBasket = ['apple', 'banana', 'orange'];
console.log(fruitBasket[0]);
console.log(fruitBasket[1]);
console.log(fruitBasket[2]);
console.log(fruitBasket.length);
fruitBasket.push('pear');
// add a new value at the beginning
fruitBasket.unshift('melon');
console.log(fruitBasket);
fruitBasket.pop();
console.log(fruitBasket);
// remove value from the beginning
fruitBasket.shift();
console.log(fruitBasket);

// Arrays are Objects!
const arr = [1, 2, 3];
console.log(typeof arr);

const obj = {prop: 'value'};
console.log(typeof obj);