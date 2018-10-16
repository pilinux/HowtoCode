/*
* Constructor:
* - a block of code similar to method.
* - it is called every time an object is created with new() keyword.
* - constructs values at the time of object creation.
* - java compiler creates a default constructor if the class doesn't have any constructor.
* - with access modifiers, private, protected, public or default constructor is created.
*
* Rules for creating Java constructor:
* - constructor name must be the same as its class name.
* - must have no explicit return type.
* - cannot be abstract, static, final, or synchronized.
*
* Two types of constructors in Java:
* - default constructor (no argument)
* -- class_name() {}
 * - parameterized constructor
*
* Note: In this example, default constructor is explained.
* */

//Java program
public class Main {

    //main method
    public static void main(String[] args) {
        //calling a default constructor
        HelloWorld helloWorld = new HelloWorld();
    }

}

class HelloWorld {
    //default constructor
    HelloWorld() {
        System.out.println("Hello World!");
    }
}


/*
* Output:
* >>> Hello World!
* */

/*
* Info:
* - if there is no constructor in a class, compiler automatically creates a default constructor.
* class HelloWorld { } --> compiler --> class HelloWorld { HelloWorld() {} }
* */
