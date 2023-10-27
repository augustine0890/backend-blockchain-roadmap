/**
 * Returns a string with the first letter of each word capialized.
 * @param {string} string 
 * @returns {string}
 */
function titleCase(string) {
    let words = string.toLowerCase().split(' ');
    let newWords = '';
    for (const word of words) {
        newWords += word[0].toUpperCase() + word.slice(1) + ' ';
    }
    return newWords;
}

function titleCase(string) {
    return string.toLowerCase().split(' ').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ');
}

console.log(titleCase("I'm a little tea pot")); // I'm A Little Tea Pot
console.log(titleCase('sHoRt AnD sToUt')); // Short And Stout
console.log(titleCase('HERE IS MY HANDLE HERE IS MY SPOUT')); // Here Is My Handle Here Is My Spout