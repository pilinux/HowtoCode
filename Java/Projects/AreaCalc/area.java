    public static double area(double radius) {

        final double PI = 3.14159;  //Constant

        if (radius < 0)
            return -1.0;

        return radius * radius * PI;
    }

    public static double area(double x, double y) {
        if (x < 0 || y < 0)
            return -1.0;

        return x * y;
    }
