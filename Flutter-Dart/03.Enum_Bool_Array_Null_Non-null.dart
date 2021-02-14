// enumerated types
// predefined list of values
// or, containers for constant values
enum Brands { BMW, Audi, Tesla }

void main() {
  // enum
  var liked = Brands.Tesla;
  print(liked); // Brands.Tesla
  print(liked.toString()); // Brands.Tesla
  print("");

  // bool
  bool test1 = 10 == 1;
  bool test2 = 10 == 10;
  print(test1); // false
  print(test2); // true
  print("");

  // array - ordered, indexable collection of objects
  // represented by a List<T>
  // List literals like JavaScript
  var list = [1, 2, 3];
  print(list is List); // true
  print(list); // [1, 2, 3]
  print(list[1]); // 2
  print(list.length); // 3
  list.add(4);
  print(list.length); // 4
  print(list); // [1, 2, 3, 4]
  list.removeAt(1);
  print(list); // [1, 3, 4]
  print(list.isEmpty); // false
  print(list.contains(2)); // false
  print("");

  // nullable and non-nullable types
  // starting from Dart 2.10, variables are
  // non-nullable by default
  // non-nullability:
  // - adds type-safety
  // - reduces runtime exceptions related to `null`
  // - no need to check `null` value
  // example (no need to do the following):
  // int value = 10;
  // void main() {
  //   if (value != null) {
  //     print(value);
  // }

  // Dart 2.10+, it won't compile
  // int value;
  // print("$value");

  int value1 = 1;
  print("$value1"); // 1

  int value2;
  value2 = 2;
  print("$value2"); // 2
  print("");

  // - usage of `null` is allowed
  // - type-safety degree is lost
  // - append a `?` at the end of the type
  int? value3;
  print("$value3"); // null
  print("");

  // convert nullable to non-nullable by adding `!`
  int? value4 = 200;
  int value5 =
      value4!; // Warning: Operand of null-aware operation '!' has type 'int' which excludes null
  print("$value5"); // 200
  value5 = 5;
  print("$value5"); // 5
  print("");
  // exception is thrown if the nullable value is actually `null`
  // int? value4;
  // int value5 = value4!; // exception is thrown

  // convert nullable to non-nullable `subtype`
  // by the typecast operator `as`
  num? value6 = 100;
  print("$value6"); // 100
  int value7 = value6 as int; // value of `num` is int type
  print("$value7"); // 100

  // better: use `??` operator to produce non-nullable value from nullable one
  int? value8;
  print("$value8");
  int value9 = value8 ?? 9;
  print("$value9");
  print("");

  // null-aware member access operator `?.`
  double? pi = 3.14159;
  final r = pi?.round();
  print("$r");
}
