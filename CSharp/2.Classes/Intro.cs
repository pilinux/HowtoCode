/*
 * Class:
 * - A group of related methods and variables
 * - Everything in C# is built upon classes
 */

using System;

namespace Intro
{
    class Program
    {
        static void Main(string[] args)
        {
            //Step 5: Variable and instance
            /*
             * - Step 5A: Declare a new variable:
             *            Novice novice;
             * - Step 5B: Create a new instance:
             *            novice = new Novice();
             * In one line:
             *            Novice novice = new Novice();
             */
            Novice novice = new Novice();
            //At step 3 a constructor was created which prints only a statement.
            //Hence at step 5, it will print that statement.
            
            //Step 4 is required to pass a string or any other type of value.
            novice = new Novice("boy & girl");
            
            //From step 2B, it uses 'getter' method.
            Console.WriteLine(novice.Gender);    //'get' is used
            
            //From step 2B, it uses 'setter' method.
            Console.WriteLine(novice.Gender = "girl");    //'set' is used
            
            //It's calling the instance and is executing that method.
            Console.WriteLine(novice.Describe());
        }
    }
    
    //Step 1: Define a new class Novice
    /*
     * class Novice {}
     */
    class Novice
    {
        //Step 2: Define class properties
        /*
         * class properties:
         * - Concept: variable + method
         * - Control the accessibility of a class variables
         * - Recommended way to access variables from outside
         * - Consists of 2 parts - a get method and a set method
         * - Only one method (get or set) is required, but two methods can be used
         * - Only get method: read-only
         * - Only set method: write-only
         * - get and set: r/w
         */
        
        //Step 2A: Define variable
        private string gender;
        
        //Step 2B: Define method (getter-setter)
        public string Gender
        {
            get => gender.ToUpper();    //return the value
            set => gender = value;    //assign a value
            //another way to write 'set'
            /*
            set
            {
                if (value == "girl")
                {
                    gender = value;
                }
            }
            */
        }
        
        //Step 3: Special methods
        /*
         * Constructor:
         * - Used to instantiate a class
         * - Never returns anything => do not have to define any return type
         */
        public Novice()
        {
            Console.WriteLine("A constructor with no parameter.");
        }
        
        //Step 4: Special method to pass variable
        /*
         * It is ok to have several constructors.
         * This constructor is used to pass string as a variable.
         * Example: novice = new Novice("boy");
         */
        public Novice(string gender)
        {
            this.gender = gender;
        }
        
        //A quick example on how to create an instance.
        public string Describe()
        {
            return "This apprentice is a " + gender;
        }
    }
}
