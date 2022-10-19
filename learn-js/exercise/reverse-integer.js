const num = 3849;

function reverseGivenInteger(num) {
  let reverseNum = 0
  let lastDigit
  while (num !== 0) {
    lastDigit = num % 10
    reverseNum = reverseNum*10 + lastDigit
    num = Math.floor(num / 10)
  }

  return reverseNum;
}

console.log(`Reversed integer is: ${reverseGivenInteger(num)}`)
console.log(`Reversed integer is: ${reverseGivenInteger(1111)}`)
console.log(`Reversed integer is: ${reverseGivenInteger(64684)}`)
console.log(`Reversed integer is: ${reverseGivenInteger(86173)}`)