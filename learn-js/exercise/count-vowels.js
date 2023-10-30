function countVowels(str) {
    const strLower = str.toLowerCase();
    const vowels = new Set(['a', 'e', 'i', 'o', 'u']);
    const count = [...strLower].filter(char => vowels.has(char)).length;
    return count
}

console.log(countVowels('Hello, World!'));
console.log(countVowels('JavaScript'))
console.log(countVowels('OpenAI Chatbot'));
console.log(countVowels('Coding Challenge'));