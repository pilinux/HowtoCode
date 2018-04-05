/*
 * Constructor:
 * - Special methods
 * - Objects are created from classes by constructor => instantiation/construction of a class
 *   [Objects are instances of a class]
 * - Can never return anything
 * - If constructor is not defined in C#, it automatically is created behind the scene to
 *   reserve memory block for the objects of the class
 */
using System;

namespace Constructor
{
    class Program
    {
        static void Main(string[] args)
        {
            //Step 6: A new object 'ana' is created from the class 'Person'.
            //It takes the default values of the variables as parameters.
            Person ana = new Person();
            ana.sayHello();
         
            //Step 7: A new object 'peter' is created from the class 'Person'.
            //"Peter" is passed as a string parameter.
            Person peter = new Person("Peter");
            peter.sayHello();
            
            //Step 8: A new object 'jack' is created from the class 'Person'.
            //"Jack" is passed as a string parameter.
            //50 is passed as an int parameter.
            Person jack = new Person("Jack", 50);
            jack.sayHello();
            
            //Step 9: A new object 'sarah' is created from the class 'Person'
            //"Sarah" is passed as a string _name parameter.
            //22 is passed as an int _age parameter.
            //155 is passed as a double _height parameter.
            Person sarah = new Person("Sarah", 22, 155);
            sarah.sayHello();
            
            
            //Step 11 (a different approach): Assigning new variable values using 'setter' method
            peter.setName("John");
            peter.setAge(30);
            peter.setHeight(170);
            peter.sayHello();
        }
    }

    //Step 1: Class
    public class Person
    {
        //Step 2: Variables with default values (or without any value)
        private string _name = "Ana";
        //'_name': it is a variable of type 'string' with a default value "Ana"
        private int _age = 20;
        //'_age': it is a variable of type 'int' with a default value 20
        private double _height = 160;
        //'_height': it is a variable of type 'double' with a default value 160

        //Step 3: Constructor
        public Person()
        {
            
        }
        
        //Step 4 (optional): Constructor can be overloaded
        public Person(string _name)    //here same variable '_name' is used as a parameter
        {
            this._name = _name;
            //since '_name' is used as the parameter, 'this' qualifier is recommended
            //format: <qualifier>.<already defined variable names at the beginning> = <parameter>
            //        this._name = _name;
        }
        public Person(string _name, int _age)
        {
            this._name = _name;
            this._age = _age;
        }
        public Person(string _name, int _age, int h)    //different variable 'h' is used as a parameter
        {
            this._name = _name;
            this._age = _age;
            _height = h;
            //'this' qualifier is not required since a different variable 'h' is used
            //format: <already defined variable names at the beginning> = <parameter>
            //        _height = h;
        }
        
        //Step 5: Method
        public void sayHello()
        {
            //variables (already defined at the beginning of this class) are used here
            Console.WriteLine("Hello, my name is " + _name);    //variable: _name
            Console.WriteLine("I'm " + _age + " years old.");    //variable: _age
            Console.WriteLine("And I'm " + _height + " cm tall.");    //variable: _height
            Console.WriteLine("");
        }

        
        //Step 10 (a different approach): Assigning new values to the variables using 'setter' method
        public void setName(string _name)
        {
            this._name = _name;
        }
        public void setAge(int a)
        {
            _age = a;
        }
        public void setHeight(int _height)
        {
            this._height = _height;
        }
    }
}
