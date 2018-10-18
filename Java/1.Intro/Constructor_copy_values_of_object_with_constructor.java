/*
* Copy constructor using java constructor:
* - there is no copy constructor in java
* - however, it is possible to copy the values from one object to another object by,
* -- constructor
* -- assigning the values of one object into another
* -- clone() method of Object class
* */

//Java program
public class Main {

    //main method
    public static void main(String[] args) {
        //creating an object f1 by calling the constructor Food
        Food f1 = new Food(1, "Maracuja");
        //copying values from f1 to f2 by using the constructor Food
        Food f2 = new Food(f1);

        f1.displayOnTerminal();
        /*
        * Output:
        * >>> Food ID is 1 and name is Maracuja
        * */
        f2.displayOnTerminal();
        /*
         * Output:
         * >>> Food ID is 1 and name is Maracuja
         * */
    }

}

class Food {
    int ID;
    String name;

    //constructor to initialize int and String
    Food(int i, String n) {
        ID = i;
        name = n;
    }

    //constructor to initialize another object
    //here constructor is copying values
    Food(Food f) {
        ID = f.ID;
        name = f.name;
    }

    void displayOnTerminal() {
        System.out.println("Food ID is " + ID + " and name is " + name);
    }
}
