using System;

namespace Array
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Array of strings:");
            string[] names = new string[2];
            //Instantiate string array variable
            
            names[0] = "John Doe";
            names[1] = "Michael Corleone";
            foreach (string name in names)
            {
                Console.WriteLine(name);
            }
            
            Console.WriteLine("");
            Console.WriteLine("Array of int Method I:");
            //Method I to define an array
            //In one line an array with a size of 5 is created
            //and filled with all items
            int[] numbersI = new int[5] { 0, 1, 2, 3, 4 };
            foreach (int num in numbersI)
            {
                Console.WriteLine(num);
            }
            
            Console.WriteLine("");
            Console.WriteLine("Array of int Method II:");
            //Method II to define an array
            int[] numbersII = { 0, 1, 2 };
            foreach (int num in numbersII)
            {
                Console.WriteLine(num);
            }
            
            //Find length of 'names' array using Length property of array
            Console.WriteLine("");
            Console.WriteLine("Length of array: " + names.Length);
            Console.WriteLine("");
            
            //Print array items with item number
            for (int i = 0; i < names.Length; i++)
            {
                Console.WriteLine("Item number " + i + ": " + names[i]);
            }
            
            Console.WriteLine("");
            int[] numbers = {40, 22, 52, 5, 0, 78, 202, -1};
            Console.WriteLine("Items in an array: ");
            foreach (int i in numbers)
            {
                Console.WriteLine(i);
            }
            
            //Reversing the order
            Console.WriteLine("");
            Console.WriteLine("In reverse order: ");
            System.Array.Reverse(numbers);
            foreach (int i in numbers)
            {
                Console.WriteLine(i);
            }
            
            //Sorting an array
            Console.WriteLine("");
            System.Array.Sort(numbers);
            Console.WriteLine("Sort the items: ");
            foreach (int i in numbers)
            {
                Console.WriteLine(i);
            }
            
            //Reversing the order
            Console.WriteLine("");
            System.Array.Reverse(numbers);
            Console.WriteLine("Again in reverse order: ");
            foreach (int i in numbers)
            {
                Console.WriteLine(i);
            }
        }
    }
}
