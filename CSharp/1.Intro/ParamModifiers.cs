/*
 * Params Modifiers: takes an arbitrary number of parameters.
 *
 * Examples in this program:
 * - FavFruit()      : Use an array as a parameter
 * - VisitCities()   : 2 int arguments + 1 char argument + unlimited parameters.
 *                     Parameter with params keyword always has to be the last one.
 *                     In this example: params string[] charParameters <= is the last one.
 * - VisitCountries(): 1 int argument + unlimited parameters.
 *
 * Methods with params can take any other arguments. Params can be passed
 * or not, but all arguments must be passed.
 */

using System;

namespace ParamModifiers
{
    class Program
    {
        static void Main(string[] args)
        {
            /*
             * 
             */
            Console.WriteLine("Favorite fruits: ");
            FavFruit(new string[] {"Mango", "Apple", "Orange"});
            
            Console.WriteLine("");
            VisitCities(1, 2, '0', "Paris", "Berlin", "10");    //args and params are passed
            VisitCountries(0);    //only arg is passed
            VisitCountries(1, "Belgium", "Netherlands");    //arg and params are passed
        }

        private static void FavFruit(string[] fruitNames)
        {
            foreach (var NAME in fruitNames)
            {
                Console.WriteLine(NAME);
            }
        }

        private static void VisitCities(int intArgumentI, int intArgumentII, char charArgument, 
            params string[] charParameters)
        {
            Console.WriteLine("Char parameters: ");
            foreach (var NAME in charParameters)
            {
                Console.WriteLine("I visited " + NAME);
            }

            Console.WriteLine("");
            Console.WriteLine("Int ArgumentI : " + intArgumentI);
            Console.WriteLine("Int ArgumentII: " + intArgumentII);
            Console.WriteLine("Char ArgumentI: " + charArgument);
            Console.WriteLine("");
        }

        private static void VisitCountries(int unusedArgument, params string[] countries)
        {
            Console.WriteLine("One argument and/or parameters: ");
            foreach (var NAME in countries)
            {
                Console.WriteLine("I visited " + NAME);
                /*
                 * For VisitCountries(0); nothing will be printed out except the default msg:
                 * One argument and/or parameters:
                 *
                 * For VisitCountries(1, "Belgium", "Netherlands"); only the 'params string[] countries'
                 * will be printed out on the screen.
                 */
            }
        }
    }
}
