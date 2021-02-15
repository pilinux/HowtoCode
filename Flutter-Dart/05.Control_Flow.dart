// enum for switch statement
enum PaymentStatus { Finished, Failed, Canceled }

void main() {
  // if...else
  int number = 14;

  if (number % 2 == 0) {
    print("$number is an even number.");
  } else {
    print("$number is an odd number.");
  }
  print("");

  // switch statement
  final paymentStatus = PaymentStatus.Finished;

  switch (paymentStatus) {
    case PaymentStatus.Finished:
      print("Payment is successful!");
      break;
    case PaymentStatus.Failed:
      print("Payment is failed!");
      break;
    case PaymentStatus.Canceled:
      print("Payment is canceled!");
      break;
    default:
      print("System error!");
  }
  print("");

  // for loop
  for (var i = 0; i < 5; i++) print("Number is $i");
  print("");

  // while loop
  var i = 0;
  while (i < 5) {
    print("Number is $i");
    i++;
  }
  print("");

  // do...while
  var j = 0;
  do {
    print("Number is $j");
    j++;
  } while (j < 5);
  print("");

  // `break` stops the loop in which it is called
  for (var k = 0; k < 5; k++) {
    for (var l = 0; l < 5; l++) {
      if (l == 3) break;
      print("k = $k");
      print("l = $l");
    }
  }
  print("");

  // for-in loop
  final List<String> brands = ["BMW", "Tesla", "Audi", "Ford"];
  for (var m = 0; m < brands.length; m++) print(brands[m]);
  print("");
  for (final brand in brands) print(brand);

  // assertion
  // - throw an exception if the given condition evaluates to false
  // - works ONLY in debug mode
  // - in production mode, compiler will ignore assertion
  // example:
  // final int w = -10;
  // assert(w < 0, "Number cannot be negative!");
  // print("");
}
