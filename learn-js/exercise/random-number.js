function randomNumberGeneratorInRange(rangeStart, rangeEnd) {
  return Math.floor(Math.random() * (rangeEnd - rangeStart + 1)) + rangeStart;
}

console.log(`My random number: ${randomNumberGeneratorInRange(5, 100)}`);
console.log(`My random number: ${randomNumberGeneratorInRange(100, 200)}`);
