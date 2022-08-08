var sum = function(x, y) {
  return x + y;
}

var onePlusOne = sum(1, 2);
var twoPlusTwenty = sum(2, 20);
console.log(onePlusOne + twoPlusTwenty);

var subtract = function(x, y) {
  return x - y;
}
console.log(subtract(1, 11));
console.log(subtract(11, 1));

var weather = function(temperature) {
  console.log("The temperature outside is", temperature, "degrees farenheight.");
  if (temperature <= 30) {
    console.log("It's freezing outside! It'll be best to bundle up.");
  } else if(temperature <= 55) {
    console.log("It's getting cold outside. Better wear a jacket!");
  } else if (temperature <= 75){
    console.log("It's pleasant outside!");  
  } else {
    console.log("It's getting hot outside!");
  }
}

weather(20);
weather(40);
weather(60);
weather(90);

var sellStock = function(stockPrice, sellPrice){
  if (stockPrice >= sellPrice) {
    return true;
  } else {
    return false;
  }
}

var num = 0;
while (num <= 10) {
  if (num !== 10) {
    console.log("The number is", num, "- less than 10");
  } else {
    console.log("The number is", num, "- the loop is now over");
  }
  num++;
}

for (var i=0; i<10; i++) {
  console.log(i);
}

var rangeSum = function(number) {
  var sum = 0;
  var i = number;
  if (number < 0) {
    while (i < 0) {
      sum += i;
      i++
    }
  } else {
    while (i > 0) {
      sum += i;
      i--;
    }
  }
  return sum;
}
console.log("Range Sum:", rangeSum(5));
console.log("Range Sum:", rangeSum(-30));