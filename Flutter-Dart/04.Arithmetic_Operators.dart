void main() {
  // add
  print(2 + 3); // 5

  // subtract
  print(5 - 2); // 3

  // multiply
  print(6 * 2); // 12

  // divide
  print(7 / 2); // 3.5

  // integer division
  print(7 ~/ 2); // 3

  // remainder
  print(7 % 2); // 1

  // increment/decrement
  int a = 10;
  // value is incremented, then returned
  print(++a); // 11
  // value is returned, then incremented
  print(a++); // 11 during printing, final value is 12
  print(--a); // 11
  print(a--); // 11 during printing, final value is 10

  int b = 20;
  b += 5;
  print(b); // 25
}
