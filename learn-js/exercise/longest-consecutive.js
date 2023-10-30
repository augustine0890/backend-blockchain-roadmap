function longestConsecutiveSequence(nums) {
    if (nums.length === 0) {
        return 0;
    }

    // Create a Set to store the numbers
    const numSet = new Set(nums);
    let longestSteak = 0;

    for (let num of numSet) {
        // If the current number is the start of a sequence
        if (!numSet.has(num - 1)) {
            // Try to extend the sequence
            let currentNum = num;
            let currentSteak = 1;

            while (numSet.has(currentNum + 1)) {
                currentNum += 1;
                currentSteak += 1;
            }

            // Update the longest streak if necessary
            longestSteak = Math.max(currentSteak, longestSteak);
        }
    }
    return longestSteak;
}

console.log(longestConsecutiveSequence([100, 4, 200, 1, 3, 2])); // Output: 4 (The longest consecutive sequence is [1, 2, 3, 4])
console.log(longestConsecutiveSequence([0, 3, 7, 2, 5, 8, 4, 6, 9, 1])); // Output: 10 (The longest consecutive sequence is [0, 1, 2, 3, 4, 5, 6, 7, 8, 9])