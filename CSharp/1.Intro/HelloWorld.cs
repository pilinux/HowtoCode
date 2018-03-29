using System;
//'using' keyword imports a namespace 'System'.
//namespace is a collection of classes.

namespace HelloWorld    //user-defined namespace
{
    class Program
    //Program class is a part of HelloWorld namespace.
    //Program class is the main class for this application.
    {
        static void Main(string[] args)
        //'Main' is a method and the entry-point of this application.
        //'static' keyword -> 'Main' method can be accessible without instantiating the class.
        //Return type is void.
        //'string[] args' -> takes arguments as an array of strings
        {
            Console.WriteLine("Hello World!");
            //'Console' class is a part of 'System' namespace.

            int number1, number2;
            string firstName, lastName;
            
            number1 = 1;
            Console.Write("How old are you? ");
            number2 = int.Parse(Console.ReadLine());
            
            Console.Write("What is your first name? ");
            firstName = Console.ReadLine();
            
            Console.Write("What is your last name? ");
            lastName = Console.ReadLine();
            
            Console.WriteLine("My name is " + firstName + " " + lastName);
            Console.WriteLine(("I am not " + number1 + " but " + number2 + " years old xD"));
        }
    }
}
