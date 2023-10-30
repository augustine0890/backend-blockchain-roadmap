function removeDuplicates(arr) {
    let seen = {};
    let result = [];

    for (let item of arr) {
        if (!seen[item]) {
            seen[item] = true;
            result.push(item);
        }
    }
    return result;
}

// function removeDuplicates(arr) {
// return arr.filter((item, index) => arr.indexOf(item) === index);
// }

console.log(removeDuplicates([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])); // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
console.log(removeDuplicates([1, 1, 1, 1, 1, 1, 1, 1, 1, 1])); // [1]
console.log(removeDuplicates([1, 2, 3, 4, 5, true, 1, 'hello', 2, 3, 'hello', true])); // [1, 2, 3, 4, 5, true, 'hello']