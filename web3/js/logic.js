let number = 4;
if ((number >= 0) && (number <= 100)) {
  console.log(`${number} is between 0 and 100, both included`);
}

const numTwo = -1;
if ((numTwo < 0) || (numTwo > 100)) {
  console.log(`${numTwo} is NOT in between 0 and 100`);
}

const numThree = 60;
if (!(number > 100)) {
  console.log(`${numThree} is less than or equal to 100`);
}