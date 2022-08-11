class Developer {
  constructor(firstname, lastname) {
    this.firstname = firstname;
    this.lastname = lastname;
  }

  getName() {
    return `${this.firstname} ${this.lastname}`;
  }
}

var me = new Developer('Robin', 'Hoods');
console.log(me.getName());


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