/*
 * Static members:
 * - We create one class and every time create a new instance of the same class to use them
 *   for different purposes.
 * - Some classes cannot be instantiated, or their methods can be used without creating any object.
 *   These classes are called static.
 * - These members (variable for example) always remain the same.
 * - A non-static class can have static members. Non-static class has to be instantiated, but the
 *   static members will be used without creating any object.
 */

using System;

namespace Static
{
    class Program
    {
        static void Main(string[] args)
        {
            //Step 2:
            //Static class Rectangle and its static members (CalculateArea etc.) are used
            //without instantiating the class and without creating any object for the members.
            Console.WriteLine("Area of the rectangle is: " + Rectangle.CalculateArea(10, 5));

            //Step 4: Use static 'SumNumbers' without instantiating non-static class 'Addition'
            //and without creating any object.
            Console.WriteLine("Summation of two given numbers: " + Addition.SumNumbers(10, 5));
        }
    }

    //Step 1: A static class
    public static class Rectangle
    {
        public static int CalculateArea(int width, int height)
        {
            return width * height;
        }
    }
    
    //Step 3: A non-static class
    public class Addition
    {
        public static int SumNumbers(int x, int y)
        {
            return x + y;
        }
    }
}
