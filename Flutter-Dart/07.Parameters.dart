void main() {
  test1(b: 4, a: 6); // 6 4
  test1(b: 4); // null 4

  test2(a: 6); // 6 0
  test2(a: 6, b: 4); // 6 4

  test3(a: 6, b: 4); // 6 4
  test3(b: 4); // 0 4

  test4(6, b: 4); // 6 4

  test5(6); // 6 0
  test5(6, 4); // 6 4
}

// optional parameters
void test1({int? a, int? b}) {
  print("$a $b");
}

// optional parameters - parameter `b` with default value
void test2({int? a, int b = 0}) {
  print("$a $b");
}

// parameter `a` is optional, parameter `b` is REQUIRED
void test3({int a = 0, required int b}) {
  print("$a $b");
}

// mixing optional and non-optional parameters
// non-optional parameters always has to be placed before optional parameters
void test4(int a, {int? b}) {
  print("$a $b");
}

// optional parameters - positional parameters
// name of the parameter(s) doesn’t have
// to be written in the function call
void test5([int? a, int b = 0]) {
  print("$a $b");
}
