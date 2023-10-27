function countOccurrences(word, letter) {
    // Ensure the letter is a single character
    if (letter.length !== 1) {
        throw new Error('Letter must be a single character');
    }

    let count = 0;
    for (let i = 0; i < word.length; i++) {
        if (word.charAt(i) === letter) {
            count++;
        }
    }
    return count;
}

console.log(countOccurrences('hello', 'l')); // 2
console.log(countOccurrences('programming', 'm')); // 2
console.log(countOccurrences('banana', 'a')); // 3
