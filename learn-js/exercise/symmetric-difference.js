function symmetricDifference(arr1, arr2) {
    const set1 = new Set(arr1);
    const set2 = new Set(arr2);
    let result = [];

    for (num of set1) {
        if (!set2.has(num)) {
            result.push(num);
        }
    }

    for (num of set2) {
        if (!set1.has(num)) {
            result.push(num)
        }
    }
    return result;
}

console.log(symmetricDifference([1, 2, 3], [3, 4, 5]));
// Output: [1, 2, 4, 5]

console.log(symmetricDifference([1, 2, 2, 3, 4], [2, 3, 3, 4, 5]));
// Output: [1, 5]

console.log(symmetricDifference([1, 2, 3, 4, 5], [5, 4, 3, 2, 1]));
// Output: []

console.log(symmetricDifference([1, 2, 3], [4, 5, 6]));
// Output: [1, 2, 3, 4, 5, 6]