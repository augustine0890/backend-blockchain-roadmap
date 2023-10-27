function calculator(num1, num2, operation) {
    switch (operation) {
        case '+':
            return num1 + num2;
        case '-':
            return num1 - num2;
        case '*':
            return num1 * num2;
        case '/':
            if (num2 === 0) {
                throw new Error('Division by zero is not allowed');
            }
            return num1 / num2;
        default:
            throw new Error('Invalid operator. Supported operators are: "+", "-", "*", "/"');
    }
}

console.log(calculator(1, 2, '+'));  // 3
console.log(calculator(10, 5, '-'));  // 5
console.log(calculator(2, 2, '*'));  // 4
console.log(calculator(10, 5, '/'));  // 2
console.log(calculator(10, 5, '//'));  // 2
