void main() {
  // every object is an instance
  // of a class
  var person = Person("F1", "L1");

  print(person);
  print(person.firstName);
  print(person.lastName);

  // classic way to assign values to variables
  person.firstName = "F2";
  person.lastName = "L2";

  print(person.firstName);
  print(person.lastName);

  // assign values to variables using the cascade operator
  // - shorthand version of assigning multiple
  //   values of the same object
  // - same can be done with methods, but
  //   the returned values will be ignored
  person
    ..firstName = "F3"
    ..lastName = "L3";

  print(person.firstName);
  print(person.lastName);

  // - cascade notation is useful for calling a series
  //   of `void` methods on the same object
  // example:
  // Test()
  //   ..test1()
  //   ..test2()
  //   ..test3();
}

// dart has no method overload
// more than one function caanot have the same name
// every function name in the same class MUST be UNIQUE
class Person {
  // instance variables
  String firstName;
  String lastName;

  // constructor
  // `this` keyword refers to the current
  // instance of this class
  Person(this.firstName, this.lastName);
}
