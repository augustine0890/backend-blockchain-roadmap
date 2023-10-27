function findMaxNumber(array) {
    if (array.length === 0) {
        return undefined;
    }

    let maxNumber = array[0];
    for (const num of array) {
        if (num > maxNumber) {
            maxNumber = num;
        }
    }
    return maxNumber;
}

console.log(findMaxNumber([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])); // 10
console.log(findMaxNumber([1, 5, 3, 9, 2]));
console.log(findMaxNumber([0, -1, -5, 2]));
console.log(findMaxNumber([10, 10, 10, 10]));