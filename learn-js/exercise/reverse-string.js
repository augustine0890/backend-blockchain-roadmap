const str = "Javascript is awesome";

function reverseAString(str) {
  let revered = "";
  for (let i=str.length-1; i>=0; i--) {
    revered += str[i];
  }
  return revered;
}

console.log(`Reversed string is: ${reverseAString(str)}`)
console.log(`Reversed string is: ${reverseAString("Peter Parker is Spiderman")}`)
console.log(`Reversed string is: ${reverseAString("Augustine")}`)