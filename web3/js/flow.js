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

var numOfMonths = 14;
var ans = numOfMonths > 12 ? "Invalid" : "Valid";
console.log(ans);

var students = [ 
  ["Mary", 10], 
  ["Barbara", 11], 
  ["David", 12], 
  ["Alex", 11] 
];

let student = {
  ID: '21',
  name: 'John',
  GPA: '3.0',
};
const {ID, name, GPA} = student;
const {name: n} = student;
console.log("Is this John?", n);

var printStudents = function(students) {
  for (var i=0; i<students.length; i++) {
    console.log(students[i]);
  }
}
printStudents(students);

var arraySum = function(numbers) {
  var sums = [];
  for (var i=0; i<numbers.length; i++) {
    var sum = 0;
    for (var j=0; j<numbers[i].length; j++) {
      sum += numbers[i][j];
    }
    sums.push(sum);
  }
  return sums;
}
console.log(arraySum([[4, 5, 6, 7]])); /* returns 22 */
console.log(arraySum([[-6, 10, 0, 4]])); /* returns 8 */

student = {
  name: "Mary",
  age: 10,
  grades: [90, 80, 95]
}
for (property in student) {
  console.log(student[property]);
}

var students = [
  { 
    name: "Mary", 
    age: 10, 
    grades: [90, 88, 95]
  }, 
  {
    name: "Joseph", 
    age: 11, 
    grades: [80, 100, 90, 96]
  }
];

var getAverages = function(students) {
  var averages = [];
  for (var i=0; i<students.length; i++) {
    var gradesArray = students[i].grades;
    var sum = 0;
    for (var j=0; j<gradesArray.length; j++) {
      sum += gradesArray[j];
    }
    var average = sum/gradesArray.length;
    averages.push(average);
  }
  return averages;
}

console.log(getAverages(students));

students = [
  {ID: 1, present: true},
  {ID: 2, present: true},
  {ID: 3, present: false},
];
const presentStudents = students.filter(function(student){return student.present;});
console.log(presentStudents);

const square = (n) => n*n;
console.log("The square of nine:", square(9));