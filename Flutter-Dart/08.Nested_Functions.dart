import 'dart:math';

// declare functions inside other functions
// cannot call the inner function from outside

void main() {
  // inner() calls randomNum()
  inner(15);
}

void inner(int value) {
  // nested function
  int randomNum() => Random().nextInt(10);

  // using the nested function
  final number = value + randomNum();
  print("$number");
}
