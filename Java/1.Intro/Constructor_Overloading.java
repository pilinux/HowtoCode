/*
* Constructor Overloading:
* - more than one constructor with different parameter lists
* - each constructor performs a different task
* - are differentiated by the compiler by the number of parameters in the list and their types
* */

//Java program
public class Main {

    //main method
    public static void main(String[] args) {
        //calling a constructor with two arguments of types int and String (in order)
        Food f2 = new Food(2, "Apple");
        //calling another constructor with three arguments of types int, String and int (in order)
        Food f3 = new Food(3, "Bread", 10);

        f2.displayOnTerminal();
        /*
        * Output:
        * >>> Food ID is 2, name is Apple and sweetness level is 0
        * */

        f3.displayOnTerminal();
        /*
         * Output:
         * >>> Food ID is 2, name is Apple and sweetness level is 10
         * */
    }

}

class Food {
    int ID;
    String name;
    int sweetness;

    //constructor with two arguments
    Food(int i, String n) {
        ID = i;
        name = n;
    }

    //constructor with three arguments
    Food(int i, String n, int s) {
        ID = i;
        name = n;
        sweetness = s;
    }

    void displayOnTerminal() {
        System.out.println("Food ID is " + ID + ", name is " + name + " and sweetness level is " + sweetness);
    }
}
