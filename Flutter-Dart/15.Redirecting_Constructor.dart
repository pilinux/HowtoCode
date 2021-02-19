// - if two constructors do almost same thing,
//   redirecting constructors help avoid
//   code duplication

void main() {
  final fraction1 = Fraction(10, 3);

  final fraction2 = Fraction.oneHalf();

  final fraction3 = Fraction.oneThird();

  final fraction4 = Fraction.whole(4);

  print("fraction1: ${fraction1.result}"); // fraction1: 3.3333333333333335
  print("fraction2: ${fraction2.result}"); // fraction2: 0.5
  print("fraction3: ${fraction3.result}"); // fraction3: 0.3333333333333333
  print("fraction4: ${fraction4.result}"); // fraction4: 4.0
}

class Fraction {
  final int _numerator;
  final int _denominator;
  final double result;

  Fraction(int numerator, int denominator)
      : this._numerator = numerator,
        this._denominator = denominator,
        this.result = numerator / denominator;

  Fraction.oneHalf() : this(1, 2);

  Fraction.oneThird() : this(1, 3);

  Fraction.whole(int val) : this(val, 1);
}
