function isValue(s) {
    let stack = [];
    const brackets = {
        '(': ')',
        '{': '}',
        '[': ']',
    };

    // Time (N)
    for (char of s) {
        if (brackets[char]) {
            stack.push(char); // Space O(N)
        } else {
            if (char !== brackets[stack.pop()]) {
                return false;
            }
        }
    }
    return (!stack.length);
}

console.log(isValue("()"));
console.log(isValue("()[]{}"));
console.log(isValue("(]"));