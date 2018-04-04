/*
 * Method overloading:
 *   Add additional (optional to be more precise) functionality by adding more
 *   parameters with default values. In C# it can be achieved differently and
 *   in an easy way.
 */

using System;

namespace MethodOverloading
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Total sum of (2,5): " + Calculator.Sum(2,5));    //Sum1
            Console.WriteLine("Total sum of (2,5,7): " + Calculator.Sum(2,5,7));    //Sum2
            Console.WriteLine("Total sum of (2,5,7,8): " + Calculator.Sum(2,5,7,8));    //Sum3
            Console.WriteLine("Total sum of (11,15,17,80, 100): " + Calculator.Sum(11,15,17,80,100));    //Sum4
        }
    }

    class Calculator
    {
        //Multiple methods are created with the same name but different parameters.
        public static int Sum(int a, int b)    //Sum1
        {
            return a + b;
        }
        
        public static int Sum(int a, int b, int c)    //Sum2
        {
            return a + b + c;
        }
        
        public static int Sum(int a, int b, int c, int d)    //Sum3
        {
            return a + b + c + d;
        }
        
        public static int Sum(int a, int b, int c, int d, int e)    //Sum4
        {
            return a + b + c + d + e;
        }
    }
}
