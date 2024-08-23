// Functions in Kotlin

/*
A parameter and an argument aren't the same thing. When you define a function,
you define the parameters that are to be passed to it when the function is called.
When you call a function, you pass arguments for the parameters.
Parameters are the variables accessible to the function, such as a name variable,
while arguments are the actual values that you pass, such as the "John" string.

Warning:
Unlike in some languages, such as Java, where a function can change the value
passed into a parameter, parameters in Kotlin are immutable.
You cannot reassign the value of a parameter from within the function body.

Execute:
$ kotlinc 1.3.functions.kt -include-runtime -d functions.jar
$ java -jar functions.jar
*/

fun main() {
    birthdayGreeting()
    println()

    val greeting = birthdayGreeting2()
    println(greeting)
    println()

    println(birthdayGreeting3("John", 30))
    println()

    println(birthdayGreeting4())
    println(birthdayGreeting4("Alice"))
    println(birthdayGreeting4("Alice", 25))
    println(birthdayGreeting4(age = 25))
    println(birthdayGreeting4(name = "Alice"))
    println(birthdayGreeting4(age = 25, name = "Alice"))
}

fun birthdayGreeting() {
    println("Happy Birthday, John!")
    println("You are now 30 years old.")
}

// Function with return type
fun birthdayGreeting2(): String {
    return "Happy Birthday, John!"
}

// Function with return type and parameters
fun birthdayGreeting3(name: String, age: Int): String {
    val nameGreeting = "Happy Birthday, $name!"
    val ageGreeting = "You are now $age years old."
    return "$nameGreeting\n$ageGreeting"
}

// Function signature:
/*
The function name with its inputs (parameters) are collectively known as the function signature.
The function signature consists of everything before the return type and is shown in the following code snippet.

fun birthdayGreeting(name: String, age: Int)

The parameters, separated by commas, are sometimes called the parameter list.
*/

// Named arguments
/*
In Kotlin, you can use named arguments to specify the name of the parameter you are passing.

birthdayGreeting3(name = "John", age = 30)
birthdayGreeting3(age = 30, name = "John")
*/

// Default arguments
fun birthdayGreeting4(name: String = "John", age: Int = 30): String {
    val nameGreeting = "Happy Birthday, $name!"
    val ageGreeting = "You are now $age years old."
    return "$nameGreeting\n$ageGreeting"
}
