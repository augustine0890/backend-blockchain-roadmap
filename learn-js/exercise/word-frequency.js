function wordFrequencyCounter(str) {
    const words = str.toLowerCase().split(/\W+/).filter(word => word);
    const wordFrequency = new Map();
    for (let word of words) {
        let count = wordFrequency.get(word) || 0;
        wordFrequency.set(word, count + 1)
    }
    return wordFrequency;
}

console.log(wordFrequencyCounter('The quick brown fox jumps over the lazy dog.'));
console.log(wordFrequencyCounter('Lorem ipsum dolor sit amet, consectetur adipiscing elit.'));