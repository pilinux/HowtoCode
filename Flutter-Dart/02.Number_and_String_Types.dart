// format code: `dart format <filename>.dart`
// run program: `dart run <filename>.dart`

void main() {
  // print stdout
  print("Hello world!");
  print("");

  // variables
  // dart compiler is smart enough to
  // infer the data types
  var value = 10;
  var myName = "piLinux";
  var myDomain = 'piLinux.me';
  print(value);
  print(myName);
  print(myDomain);
  print("");

  // types can be defined explicitly
  // for better readability
  // int value = 10;
  // String myName = "piLinux";

  // `Object` and `dynamic` should never be used
  // never use:
  // dynamic value = 10;

  // `final` can be set ONLY ONCE
  // its content cannot be changed later
  final name = "Newton";
  print(name);
  // type can be used with `final`,
  // but not necessary
  final int number = 20;
  print(number);
  print("");

  // numbers
  // two data types (subclasses of `num`):
  // - `int` (max 64-bit)
  // - `double` (max 64.bit)
  int x = 10;
  double y = 5;
  num z = y - x;
  print(z); // -5.0
  print("");

  // compile-time constant
  const pi = 3.14159;
  print(pi);
  print("");

  // exponential
  var a = 1.5e2;
  print(a); // 150.0
  print("");

  // hexadecimal
  var b = 0xF1C;
  print(b); // 3868
  print("");

  // parsing numbers
  // note: parse() is deprecated
  var c = int.tryParse("120.1");
  print(c); // null
  var d = double.tryParse("120.1");
  print(d); // 120.1
  print("");

  // numbers to strings
  String x1 = 10.toString();
  print(x1); // 10
  String x2 = 10.14545.toString();
  print(x2); // 10.14545
  String x3 = 10.14545.toStringAsFixed(2);
  print(x3); // 10.15
  print("");

  // interpolate integer into a string
  var myAge = 12;
  var myStatement = "I'm $myAge years old.";
  print(myStatement);
  myStatement = "I'm ${myAge.abs()} years old.";
  print(myStatement);
  print("");

  // avoid redundancy:
  // do not do this!
  // because ${} calls toString()
  // myStatement = "I'm ${myAge.toString()} years old";

  // multiline string
  // useful for SQL queries
  var query = """
    SELECT *
    FROM posts;
  """;
  print(query);
  print("");

  // dart doesn't have `char` type
  // to access a character of a string,
  // use `[]` operator
  print(name[0]); // N
  print(name[1]); // e
  print(name[2]); // w
  print(name[3]); // t
  print(name[4]); // o
  print(name[5]); // n
  print("");

  // concatenate strings
  var s1 = 'I am ' + name + ' and I am ' + (12).toString() + ' years old.';
  var s2 = 'I am $name and I am $myAge years old.';
  var s3 = 'I am'
      ' $name '
      'and I am'
      ' $myAge '
      'years old.';
  print(s1);
  print(s2);
  print(s3);
}
