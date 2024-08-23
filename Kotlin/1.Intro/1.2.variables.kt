// Variables in Kotlin

/*
- In Kotlin, it's recommended to use the `val` keyword over the `var` keyword when possible.
- The 'val' keyword is used to declare a variable that cannot be changed.
- The 'var' keyword is used to declare a variable that can be changed.

Note on type inference:
- Kotlin is a statically typed language, but it also supports type inference.
- Type inference allows the compiler to determine the type of a variable based on the value assigned to it.
- Example: val name = "John" // The compiler will infer the type of 'name' as String.

Execute:
$ kotlinc 1.2.variables.kt -include-runtime -d variables.jar
$ java -jar variables.jar
*/

fun main() {
    // Declaring a variable using the 'val' keyword
    val name: String = "John"
    println("Name: $name")

    // Declaring a variable using the 'var' keyword
    var age: Int = 25
    println("Age: $age")

    // Changing the value of a variable
    age = 30
    println("Age: $age")

    // Type inference
    val country = "USA"
    println("Country: $country")

    // Double
    val pi = 3.14
    println("PI: $pi")

    // Float
    val weight = 75.5f
    println("Weight: $weight")

    // Boolean
    val isStudent = true
    println("Is student: $isStudent")

    // Character
    println("Say \"hello\"")
}
