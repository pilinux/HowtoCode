// - generally used to implement a default behavior the user expects
// - alternative to have multiple constructors
//   as dart has no method overload

void main() {
  final fraction1 = Fraction(10, 3);

  // using a named constructor
  final fraction2 = Fraction.zero();

  print("fraction1: ${fraction1.result}"); // fraction1: 3.3333333333333335
  print("fraction2: ${fraction2.result}"); // fraction2: 0.0
}

class Fraction {
  final int _numerator;
  final int _denominator;
  final double result;

  Fraction(int numerator, int denominator)
      : this._numerator = numerator,
        this._denominator = denominator,
        this.result = numerator / denominator;

// denominator cannot be zero
  Fraction.zero()
      : this._numerator = 0,
        this._denominator = 1,
        this.result = 0;
}
