function anagramGrouping(words) {
    const anagramGroups = new Map();
    for (let word of words) {
        const sortedChars = word.split('').sort().join(',');
        anagramGroups.has(sortedChars) ? anagramGroups.get(sortedChars).push(word) : anagramGroups.set(sortedChars, [word]);
    }
    return Array.from(anagramGroups.values());
}

console.log(anagramGrouping(['cat', 'act', 'dog', 'god', 'tac']));
// Output: [['cat', 'act', 'tac'], ['dog', 'god']]

console.log(anagramGrouping(['listen', 'silent', 'enlist', 'hello', 'world']));