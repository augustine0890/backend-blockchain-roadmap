function phoneNumberDirectory(array) {
    const phoneNumberMap = new Map();
    array.forEach(entry => {
        const [name, number] = entry.split(':');
        phoneNumberMap.set(name, number);
    });
    return phoneNumberMap;
}

const phoneNumbers = [
    'John:123-456-7890',
    'Jane:987-654-3210',
    'Joe:555-555-5555',
];

console.log(phoneNumberDirectory(phoneNumbers));