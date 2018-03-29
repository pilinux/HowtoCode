using System;
using System.Collections;

namespace ForEach
{
    class Program
    {
        static void Main(string[] args)
        {
            ArrayList list = new ArrayList();
            //'ArrayList' class is a part of 'System.Collections' namespace.
            //An instance 'list' of an ArrayList is created.
            
            //Strings are added to 'list' instance.
            list.Add("John Doe");
            list.Add("Jane Doe");
            list.Add("Someone else");
            
            //'foreach' loop runs through each item, saves them in 'NAME'
            //variable and prints them out on the console.
            foreach (string NAME in list)
            {
                Console.WriteLine(NAME);
            }
        }
    }
}
