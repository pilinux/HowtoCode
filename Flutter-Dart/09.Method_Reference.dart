void main() {
  final nums = [1, 2, 3, 4, 5];

  nums.forEach(showNum);
}

void showNum(int value) {
  print("$value");
}

// note:
// - same signature!
// - forEach() method asks for a function with a single integer parameter
//   and no return type (void)
// - showNumber() function accepts an integer as parameter
//   and returns nothing (void)
