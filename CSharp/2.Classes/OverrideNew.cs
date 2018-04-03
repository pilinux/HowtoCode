/*
 * Usage: 'override' & 'new' keywords or modifiers
 */
using System;

namespace OverrideNew
{
    class Program
    {
        static void Main(string[] args)
        {
            BaseClass baseClass = new BaseClass();
            //Variable 'baseClass' is of type 'BaseClass' and it's value is of type 'BaseClass'
            
            DerivedClass derivedClass = new DerivedClass();
            //Variable 'derivedClass' is of type 'DerivedClass' and it's value is of type 'DerivedClass'
            
            BaseClass baseClassDerivedClass = new DerivedClass();
            //Variable 'baseClassDerivedClass' is of type 'BaseClass' and it's value is of type 'DerivedClass'
            
            /*
             * Parent       : 'BaseClass'
             * ----Children : 'DerivedClass'
             *
             * Without casting, variables 'baseClass' & 'baseClassDerivedClass' have direct access only to
             * Method0, Method1 & Method2 because they are of type 'BaseClass'
             *
             * Variable 'derivedClass' can access Method0, Method1, Method2 & Method3 because it is of type
             * 'DerivedClass'
             */
            
            
            baseClass.Method0();    //Direct access to 'BaseClass'
            baseClass.Method1();    //Direct access to 'BaseClass'
            baseClass.Method2();    //Direct access to 'BaseClass'
            
            
            derivedClass.Method0();    //Inherited from 'BaseClass'
            derivedClass.Method1();
            /* 
             * 'new' modifier:
             * Inherited from 'BaseClass', hiding 'Method1' of 'BaseClass' and assigning new value to
             *  'Method1' in 'DerivedClass'
             */
            derivedClass.Method2();
            /*
             * 'override' modifier:
             * Inherited from 'BaseClass' and overriding value of 'Method2' in 'DerivedClass'
             */
            derivedClass.Method3();    //Direct access to 'DerivedClass'
            
            
            baseClassDerivedClass.Method1();    //Inherited from 'BaseClass'
            baseClassDerivedClass.Method2();    //'override' modifier
            
            /*
             * 'new' modifier cannot be used for variable 'baseClassDerivedClass'
             * since it is of type 'BaseClass'
             */
        }
    }

    class BaseClass
    {
        public void Method0()
        {
            Console.WriteLine("Base class - method0");
        }
        
        public void Method1()
        {
            Console.WriteLine("Base class - method1");
        }

        public virtual void Method2()
        {
            Console.WriteLine("Base class - method2");
        }
    }

    class DerivedClass : BaseClass    //'DerivedClass' is inherited from 'BaseClass'
    {
        public new void Method1()
        {
            Console.WriteLine("Derived class new - method1");
        }

        public override void Method2()
        {
            Console.WriteLine("Derived class overrided - method2");
        }

        public void Method3()
        {
            Console.WriteLine("Derived class (not inherited) - method3");
        }
    }
}
