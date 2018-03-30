using System;

namespace Modifier
{
    class Program
    {
        static void Main(string[] args)
        {
            int number1 = 20;
            Console.WriteLine("Initial number: " + number1);
            AddFive(number1);
            Console.WriteLine("After adding five (no modifier is used): " + number1);
            //Original value of 'number1' is not affected.
            
            RefAddFive(ref number1);
            Console.WriteLine("After adding five (ref modifier is used): " + number1);
            //Original value of 'number1' is modified.
            
            OutAddFive(out number1);
            Console.WriteLine("After adding five (out modifier is used): " + number1);
            //Original value of 'number1' is modified.
        }

        static void AddFive(int number)
        {
            number = number + 5;
        }
        
        static void RefAddFive(ref int number)
        {
            number = number + 5;
        }
        static void OutAddFive(out int number)
        {
            number = 2;
            //'out' modifiers can be thought of as an additional return variable (not input).
        }
    }
}
