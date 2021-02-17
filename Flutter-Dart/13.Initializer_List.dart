// - allows fields to be defined before the constructor body
// - required for `final` fields

void main() {
  var circle = AreaOfCircle(5);
  print("Radius: ${circle.r}"); // Radius: 5
  print("Area: ${circle.area}"); // Area: 78.53975
}

class AreaOfCircle {
  static const pi = 3.14159;
  final int r;
  final double area;

  AreaOfCircle(int r)
      : this.r = r,
        area = pi * r * r;
}
