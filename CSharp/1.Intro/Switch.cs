using System;

namespace Switch
{
    class Program
    {
        static void Main(string[] args)
        {
            int select;
            Console.Write("Select from 0 to 4: ");
            select = int.Parse(Console.ReadLine());

            switch (select)
            {
                case 0:
                    Console.WriteLine("It is number zero.");
                    break;
                case 1:
                    Console.WriteLine("It is number one.");
                    break;
                case 2:
                    Console.WriteLine("It is number two.");
                    break;
                case 3:
                    Console.WriteLine("It is number three.");
                    break;
                case 4:
                    Console.WriteLine("It is number four.");
                    break;
                default:
                    Console.WriteLine("You gave a wrong input.");
                    break;
            }
        }
    }
}
