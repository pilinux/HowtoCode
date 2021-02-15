// ALWAYS define return type
// DO NOT depend on `dynamic` return type
// use `void` when there is nothing to return

void main() {
  print(checkEven(1));
  print(checkEven(2));
  printStatement("Hello from earth!");

  // in dart, functions are objects
  // the type is called `Function`

  // assign a function to a variable
  bool Function(int) checker1 = checkEven;
  print(checker1(10));

  // typical automatic type deduction - better approach
  final checker2 = checkEven;
  var checker3 = checkEven;
  print(checker2(11));
  print(checker3(12));

  // anonymous function
  // - nameless functions created on-the-fly
  // - are immediately assigned to a variable

  // takes no parameter
  final addNumber = () => 10 + 12.4;
  print(addNumber());

  // takes parameter
  final fullname = (String nickname) {
    var myName = "Isaac";
    myName += " " + nickname;
    return myName;
  };
  print(fullname("Newton"));
  print("");

  // inside named function, parameter is passed
  // into the anonymous function
  printNum((int value) {
    print("Number $value");
  });
  printNum((int value) => print("${value + 2}"));
  print("");

  // forEach callbacks
  final anotherList = [1, 2, 3, 4, 5];
  // iterate
  anotherList.forEach((int x) => print("Number $x"));
  // note: avoid using anonymous functions in forEach() calls
}

// arrow syntax:
// - works when body of the func contains only one line
bool checkEven(int value) => value % 2 == 0;

void printStatement(String value) => print(value);

// parameter `action` accepts a `void` func with one integer
void printNum(void Function(int) action) {
  final list = [1, 2, 3, 4, 5];

  for (final item in list) {
    action(item);
  }
}
