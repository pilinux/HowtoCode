using System;

namespace Methods
{
    class Program
    {
        static void Main(string[] args)
        {
            Program sum = new Program();
            int finalResult = sum.AddNumbers(5, 7);
            //Since 'AddNumbers' method is not static, it needs to be instantiated.
            Console.WriteLine("Addition: " + finalResult);
            
            Console.WriteLine("");
            Console.Write("Multiplication:" + Multiply(5, 4.30));
            //'Multiply' method is static, so it does not need to be instantiated.
        }

        public int AddNumbers(int num1, int num2)
        {
            int result = num1 + num2;
            return result;    //Return the result
        }

        public static double Multiply(double num1, double num2)
        {
            double result = num1 * num2;
            return result;    //Return the result
        }
    }
}
