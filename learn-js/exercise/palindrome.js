function isPalindrome(str) {
    str = str.toLowerCase().replace(/[^a-z0-9]/g, '');
    let start = 0;
    let end = str.length - 1;
    while (start < end) {
        if (str.charAt(start) !== str.charAt(end)) {
            return false;
        }
        start++;
        end--;
    }
    return true;
}

console.log(isPalindrome('racecar')); // true
console.log(isPalindrome('Hello')); // false
console.log(isPalindrome('A man, a plan, a canal, Panama')); // true
console.log(isPalindrome('12321')); // true
