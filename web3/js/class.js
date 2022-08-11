class Developer {
  constructor(firstname, lastname) {
    this.firstname = firstname;
    this.lastname = lastname;
  }

  getName() {
    return `${this.firstname} ${this.lastname}`;
  }
}

var me = new Developer('Augustine', 'Nguyen');
// console.log(me.firstname);

var getName = () => console.log(this.firstname);
const printMyName = getName.bind(me);
printMyName();


var printName = function() {
  console.log(`My name is ${this.firstname} ${this.lastname}`);
};
const newPrintName = printName.bind(me);
// bound newPrintName() prints appropriately
newPrintName();
// unbound printName() prints undefined
printName();

class Car {
  constructor(color, model, engineCap, registrationNumber) {
    this.color = color;
    this.model = model;
    this.engineCap = engineCap;
    this.registrationNumber = registrationNumber;
  }

  getColor() {
    return this.color;
  }

  getModel() {
    return this.model;
  }

  setColor(color) {
    this.color = color;
  }

  setModel(model) {
    this.model = model;
  }
}

var printInfo = function(firstname, lastname, lang1, lang2, lang3) {
  this.firstname = firstname;
  this.lastname = lastname;
  console.log(`My name is ${this.firstname} ${this.lastname} and I know ${lang1}, ${lang2}, and ${lang3}`)
}

languages = ['Javascript', 'Golang', 'Rust'];
printInfo('Augustine', 'Nguyen', ...languages)

class BackendDeveloper extends Developer {
  getJob() {
    return 'Backend Developer';
  }
}
me = new BackendDeveloper('Augustine', 'Nguyen')
console.log(me.getName());
console.log(me.getJob());