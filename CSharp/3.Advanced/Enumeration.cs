/*
 * Enumerations:
 * - special sets of named values mapped to a set of integers
 * - come in handy when it is required to choose between a set of constant values,
 *   and with each possible value relating to a number
 * - enumerations can be accessed from all classes within the same namespace since
 *   they are defined above classes, but inside the namespace
 */

using System;

namespace Enumeration
{
    //Enumerations - Games & Days
    
    public enum Games { Cricket, Soccer, Tennis, Chess, Basketball }
    /**
     * Mapped values:
     * - Cricket = 0
     * - Soccer = 1
     * - Tennis = 2
     * - Chess = 3
     * - Basketball = 4
     */
    public enum Days { Monday, Tuesday = 4, Wednesday, Thursday, Friday = 8, Saturday, Sunday }
    /*
     * Mapped values:
     * - Monday = 0
     * - Tuesday = 4
     * - Wednesday = 5
     * - Thursday = 6
     * - Friday = 8
     * - Saturday = 9
     * - Sunday = 10
     *
     * When a value is mapped to an user-defined integer, all the following values will be mapped
     * to one number higher than the previous one automatically.
     */

    class Program
    {
        static void Main(string[] args)
        {
            Games basketball = Games.Basketball;
            //'basketball' object is created with 'Basketball' value from 'Games' enum
            Console.WriteLine(basketball);
            //print the value of 'basketball' on the console which is 'Basketball'
            Console.WriteLine((int) basketball);
            //print the mapped value of 'Basketball' which is 4
            Console.WriteLine();

            Games games = (Games) 3;
            //create 'games' object with the value that is mapped to integer 3 in 'Games' enum
            Console.WriteLine(games);
            //print the value of 'games' on the console which is 'Chess'
            Console.WriteLine();

            //create an array object 'gameNames' with all the values of 'Games' enum
            string[] gameNames = Enum.GetNames(typeof(Games));
            //print all values of the array
            foreach (var VARIABLE in gameNames)
            {
                Console.WriteLine(VARIABLE);
            }
            Console.WriteLine("**********");
            Console.WriteLine();
            
            
            //Another example with 'Days' enum
            Days monday = Days.Monday;
            Console.WriteLine(monday);
            Console.WriteLine((int)monday);
            
            Days tuesday = Days.Tuesday;
            Console.WriteLine(tuesday);
            Console.WriteLine((int)tuesday);
            
            Days wednesday = Days.Wednesday;
            Console.WriteLine(wednesday);
            Console.WriteLine((int)wednesday);
            
            Days thursday = Days.Thursday;
            Console.WriteLine(thursday);
            Console.WriteLine((int)thursday);
            
            Days friday = Days.Friday;
            Console.WriteLine(friday);
            Console.WriteLine((int)friday);
            
            Days saturday = Days.Saturday;
            Console.WriteLine(saturday);
            Console.WriteLine((int)saturday);
            
            Days sunday = Days.Sunday;
            Console.WriteLine(sunday);
            Console.WriteLine((int)sunday);

            Console.WriteLine();
            
            Days days = (Days) 4;
            Console.WriteLine(days);
            
            Console.WriteLine();
            
            string[] values = Enum.GetNames(typeof(Days));
            foreach (var s in values)
            {
                Console.WriteLine(s);
            }
        }
    }
}
