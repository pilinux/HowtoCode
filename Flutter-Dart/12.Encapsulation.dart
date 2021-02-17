// In dart, every member is public by default
// dart doesn't have `public`, `protected`, `private` keywords
// appending an underscore `_` in front of a member makes it
// private to its library

// only for an example:
// import 'package:customfile.dart' as custom;

void main() {
  var obj = Test();

  // OK
  var name1 = obj.firstname;
  // OK, since class `Test` resides also in this same file
  var name2 = obj._lastname;

  // just an example:
  // var obj2 = custom.AnyClass();
  // var var1 = obj2.anyName1; // OK
  // var var2 = obj2._anyName2; // ERROR, doesn't compile since only visible inside customfile.dart

  // more example:
  // var obj3 = custom._AnyMethod(); // ERROR, only visible inside customfile.dart
}

class Test {
  String firstname = "";
  String _lastname = "";
}
